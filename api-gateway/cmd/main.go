package main

import (
	"api-gateway/internal/https"
	cloudflareconnection "api-gateway/internal/pkg/CloudFlareConnection"
	"api-gateway/internal/pkg/config"
	authservice "api-gateway/internal/pkg/connectServices/auth-service"
	commentservice "api-gateway/internal/pkg/connectServices/comment-service"
	courseservice "api-gateway/internal/pkg/connectServices/course-service"
	likeservice "api-gateway/internal/pkg/connectServices/like-service"
	postservice "api-gateway/internal/pkg/connectServices/post-service"
	"api-gateway/internal/service"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	authDial, err := authservice.DialWithAuthService(cfg)
	if err != nil {
		log.Fatal(err)
	}
	postDial, err := postservice.DialWithPostService(cfg)
	if err != nil {
		log.Fatal(err)
	}
	likeDial, err := likeservice.DialWithLikeService(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	commentDial, err := commentservice.DialWithCommentService(*cfg)
	if err != nil {
		log.Fatal(err)
	}

	courseDial, err := courseservice.DialWithCourseService(*cfg)
	if err != nil {
		log.Fatal(err)
	}
	cloudFlare, err := cloudflareconnection.NewCloudflareStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	service := service.NewServiceRepositoryClient(*authDial, *postDial, *likeDial, *commentDial, *courseDial)
	r := https.Newgin(service, cloudFlare)

	target := fmt.Sprintf("%s:%d", cfg.ApiGateway.Host, cfg.ApiGateway.Port)
	log.Printf("Api Gateway Runnig on %s", target)
	if err := r.Run(target); err != nil {
		log.Fatal(err)
	}
}
