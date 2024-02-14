package main

import (
	"GroupAssist/internal/config"
	"GroupAssist/internal/repository"
	"GroupAssist/internal/service"
	"GroupAssist/internal/transport/rest"
	"GroupAssist/pkg/database"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	cache "github.com/psevdocoder/InMemoryCacheTTL"
	"log"
)

// @title		Group Assistant API
// @version		1.0
// @description	This is a sample server celler server.
// @host		localhost:8080
// @securityDefinitions.basic	BasicAuth
// @securityDefinitions.apikey Bearer Token Authentication
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(conf)
	if err != nil {
		log.Fatal(err)
	}

	defer func(db *sqlx.DB) {
		_ = db.Close()
	}(db)

	repositories := repository.InitRepositories(db)
	services := service.InitServices(repositories, conf)

	handler := rest.NewHandler(services)

	memoryCache := cache.New(conf.Cache.SearchExpiredTime)
	r := handler.Init(memoryCache, conf.Cache.TTL)

	if err = r.Run(fmt.Sprintf(":%d", conf.Server.Port)); err != nil {
		log.Fatal(err)
	}

}
