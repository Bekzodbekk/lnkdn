package main

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/postgres"
	"auth-service/internal/pkg/redis"
	RunService "auth-service/internal/pkg/service"
	"auth-service/internal/repository"
	"auth-service/internal/service"

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
	rdb := redis.InitRedis(cfg)
	queries := postgres.NewQueries(db)

	repo := repository.NewUserRepo(db, queries, rdb, cfg)
	service := service.NewService(repo)
	runService := RunService.NewRunService(service)

	log.Printf("Auth Service running on :%d port", cfg.AuthServicePort)
	if err := runService.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
