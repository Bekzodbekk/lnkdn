package cloudflareconnection

import (
	"api-gateway/internal/pkg/config"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type CloudflareStorage struct {
	client   *s3.Client
	bucket   string
	bucketId string
}

type UploadResult struct {
	URL  string `json:"url"`
	Key  string `json:"key"`
	Size int64  `json:"size"`
}

func NewCloudflareStorage(cfg *config.Config) (*CloudflareStorage, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.CloudFlare.AccountId),
		}, nil
	})

	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithEndpointResolverWithOptions(r2Resolver),
		awsconfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			cfg.CloudFlare.AccessKey,
			cfg.CloudFlare.SecretKey,
			"",
		)),
		awsconfig.WithRegion("auto"),
		awsconfig.WithRetryMaxAttempts(3), // Retry qo'shish
	)
	if err != nil {
		return nil, fmt.Errorf("AWS config error: %w", err)
	}

	client := s3.NewFromConfig(awsCfg)

	return &CloudflareStorage{
		client:   client,
		bucket:   cfg.CloudFlare.BucketName,
		bucketId: cfg.CloudFlare.BucketId,
	}, nil
}

// UploadFile - fayl yuklab, URL qaytaradi (yaxshilangan versiya)
func (s *CloudflareStorage) UploadFile(ctx context.Context, key string, file io.Reader) (*UploadResult, error) {
	// Timeout qo'shish
	uploadCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// Faylni xotiraga o'qish (size aniqlash uchun)
	var buf bytes.Buffer
	size, err := buf.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if size == 0 {
		return nil, fmt.Errorf("file is empty")
	}

	log.Printf("Uploading file: key=%s, size=%d bytes", key, size)

	// S3 ga yuklash
	putInput := &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		Body:          bytes.NewReader(buf.Bytes()),
		ContentLength: aws.Int64(size),
		// Content-Type ni avtomatik aniqlash uchun
		ContentType: aws.String(s.detectContentType(key)),
	}

	_, err = s.client.PutObject(uploadCtx, putInput)
	if err != nil {
		return nil, fmt.Errorf("S3 upload failed: %w", err)
	}

	log.Printf("File uploaded successfully: key=%s", key)

	// Natijani qaytarish
	result := &UploadResult{
		URL:  s.GenerateURL(key),
		Key:  key,
		Size: size,
	}

	return result, nil
}

// Content-Type ni aniqlash
func (s *CloudflareStorage) detectContentType(filename string) string {
	ext := ""
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			ext = filename[i:]
			break
		}
	}

	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".webp":
		return "image/webp"
	case ".pdf":
		return "application/pdf"
	case ".mp4":
		return "video/mp4"
	default:
		return "application/octet-stream"
	}
}

// UploadFileWithPublicURL - oddiy versiya (backward compatibility)
func (s *CloudflareStorage) UploadFileWithPublicURL(ctx context.Context, key string, file io.Reader) (string, error) {
	result, err := s.UploadFile(ctx, key, file)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

// DeleteFile - faylni o'chirish
func (s *CloudflareStorage) DeleteFile(ctx context.Context, key string) error {
	deleteCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err := s.client.DeleteObject(deleteCtx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	log.Printf("File deleted successfully: key=%s", key)
	return nil
}

// FileExists - fayl mavjudligini tekshirish
func (s *CloudflareStorage) FileExists(ctx context.Context, key string) (bool, error) {
	checkCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	_, err := s.client.HeadObject(checkCtx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		// NoSuchKey error - fayl mavjud emas
		return false, nil
	}

	return true, nil
}

// GenerateURL - fayl uchun public URL yaratish
func (s *CloudflareStorage) GenerateURL(key string) string {
	return fmt.Sprintf("https://pub-%s.r2.dev/%s", s.bucketId, key)
}

// GetFileInfo - fayl haqida ma'lumot olish
func (s *CloudflareStorage) GetFileInfo(ctx context.Context, key string) (*s3.HeadObjectOutput, error) {
	infoCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return s.client.HeadObject(infoCtx, &s3.HeadObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
}
