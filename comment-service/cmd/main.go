package main

import (
	"comment-service/internal/pkg/config"
	"comment-service/internal/pkg/postgres"
	RunService "comment-service/internal/pkg/service"
	"comment-service/internal/repository"
	"comment-service/internal/service"
	"log"
)

func main() {
	cfg, err := config.LOAD("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	queries := postgres.NewQueries(db)

	repo := repository.NewCommentRepo(db, queries)
	service := service.NewService(repo)
	runService := RunService.NewRunService(service)

	log.Printf("Comment service running on :%d port", cfg.CommentServicePort)
	if err := runService.RUN(cfg); err != nil {
		log.Fatal(err)
	}
}
