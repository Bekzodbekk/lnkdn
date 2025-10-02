package main

import (
	"like-service/internal/pkg/config"
	"like-service/internal/pkg/postgres"
	RunService "like-service/internal/pkg/service"
	"like-service/internal/repository"
	"like-service/internal/service"
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
	repo := repository.NewLikeRepo(queries)
	service := service.NewService(repo)
	runService := RunService.NewRunService(service)

	log.Printf("Like Service Running on :%d port", cfg.LikeServicePort)
	if err := runService.RUN(cfg); err != nil {
		log.Fatal(err)
	}

}
