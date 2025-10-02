package main

import (
	"course-service/internal/pkg/config"
	"course-service/internal/pkg/postgres"
	RunService "course-service/internal/pkg/service"
	"course-service/internal/repository"
	"course-service/internal/service"
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
	repo := repository.NewCourseRepo(db, queries)
	service := service.NewService(repo)
	runService := RunService.NewRunService(service)

	log.Printf("Course Service Running on :%d port", cfg.CourseServicePort)
	if err := runService.RUN(*cfg); err != nil {
		log.Fatal(err)
	}
}
