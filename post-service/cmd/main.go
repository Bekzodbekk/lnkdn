package main

import (
	"log"
	"post-service/internal/pkg/config"
	"post-service/internal/pkg/postgres"
	RunService "post-service/internal/pkg/service"
	"post-service/internal/repository"
	"post-service/internal/service"
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
	repo := repository.NewPostRepo(db, queries)
	service := service.NewService(repo)
	runService := RunService.NewRunService(service)

	log.Printf("Post Service running on :%d port", cfg.PostServicePort)
	if err := runService.RUN(cfg); err != nil {
		log.Fatal(err)
	}
}
