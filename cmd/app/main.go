package main

import (
	"GroupAssist/internal/config"
	"GroupAssist/internal/repository"
	"GroupAssist/internal/service"
	"GroupAssist/internal/transport/rest"
	"GroupAssist/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	cache "github.com/psevdocoder/InMemoryCacheTTL"
	"log"
)

// @title			Swagger Group Assistant API
// @version		1.0
// @description	This is a sample server celler server.
// @host		localhost:8080
// @securityDefinitions.basic	BasicAuth
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	pgConf, err := config.InitPostgres()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(pgConf)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	repositories := repository.InitRepositories(db)
	services := service.InitServices(repositories)

	handler := rest.NewHandler(services)

	cacheConf, err := config.InitCache()
	if err != nil {
		log.Fatal(err)
	}

	memoryCache := cache.New(cacheConf.SearchExpiredTime)

	r := handler.Init(memoryCache, cacheConf.TTL)

	handler.InitAPI(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
