package main

import (
	"log"
	"ozinshe"
	"ozinshe/pkg/handler"
	"ozinshe/pkg/repository"
	"ozinshe/pkg/service"

	_ "github.com/lib/pq"
)

func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		// Host:     "localhost",
		Host:     "db",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "ozinshe",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := ozinshe.Server{}
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("произошла ошибка при запуске http-сервера: %s", err.Error())
	}
}
