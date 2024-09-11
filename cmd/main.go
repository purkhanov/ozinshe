package main

import (
	"log"
	"ozinshe"
	_ "ozinshe/docs"
	"ozinshe/pkg/handler"
	"ozinshe/pkg/repository"
	"ozinshe/pkg/service"

	_ "github.com/lib/pq"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title           Ozinshe
// @version         1.0.0

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /api/v1
// @schemes http https

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db, err := repository.NewPostgresDB(repository.Config{
		Host: "localhost",
		// Host:     "db",
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
