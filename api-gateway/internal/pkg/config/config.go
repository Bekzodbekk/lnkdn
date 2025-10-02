package config

import (
	"github.com/spf13/viper"
)

type AuthService struct {
	Host string
	Port int
}
type PostService struct {
	Host string
	Port int
}
type LikeService struct {
	Host string
	Port int
}
type CommentService struct {
	Host string
	Port int
}
type CourseService struct {
	Host string
	Port int
}
type ApiGateway struct {
	Host string
	Port int
}
type CloudFlare struct {
	AccountId  string
	AccessKey  string
	SecretKey  string
	BucketName string
	BucketId   string
}
type Config struct {
	AuthService    AuthService
	PostService    PostService
	LikeService    LikeService
	CommentService CommentService
	CourseService  CourseService

	ApiGateway ApiGateway

	CloudFlare CloudFlare
	JWTToken   string
}

func LOAD(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := Config{
		AuthService: AuthService{
			Host: viper.GetString("authservice.host"),
			Port: viper.GetInt("authservice.port"),
		},
		PostService: PostService{
			Host: viper.GetString("postservice.host"),
			Port: viper.GetInt("postservice.port"),
		},
		LikeService: LikeService{
			Host: viper.GetString("likeservice.host"),
			Port: viper.GetInt("likeservice.port"),
		},
		CommentService: CommentService{
			Host: viper.GetString("commentservice.host"),
			Port: viper.GetInt("commentservice.port"),
		},
		CourseService: CourseService{
			Host: viper.GetString("courseservice.host"),
			Port: viper.GetInt("courseservice.port"),
		},
		ApiGateway: ApiGateway{
			Host: viper.GetString("apigateway.host"),
			Port: viper.GetInt("apigateway.port"),
		},
		CloudFlare: CloudFlare{
			AccountId:  viper.GetString("cloud_flare.account_id"),
			AccessKey:  viper.GetString("cloud_flare.access_key"),
			SecretKey:  viper.GetString("cloud_flare.secret_key"),
			BucketName: viper.GetString("cloud_flare.bucket_name"),
			BucketId:   viper.GetString("cloud_flare.bucket_id"),
		},

		JWTToken: viper.GetString("secret-key"),
	}

	return &cfg, nil
}
