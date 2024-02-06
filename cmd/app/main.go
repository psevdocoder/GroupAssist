package main

import (
	"GroupAssist/internal/config"
	"GroupAssist/internal/repository"
	"GroupAssist/internal/service"
	"GroupAssist/internal/transport/rest"
	"GroupAssist/pkg/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

//	@title			Swagger Group Assistant API
//	@version		1.0
//	@description	This is a sample server celler server.

//	@host		localhost:8080

//	@securityDefinitions.basic	BasicAuth

func main() {
	pgConf := config.InitPostgres()
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

	repositories := repository.InitRepositoties(db)
	services := service.InitServices(repositories)

	//sub, err := services.Subject.getSubjectByID(1)
	//
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(sub)

	handler := rest.NewHandler(services)

	r := handler.Init()

	handler.InitAPI(r)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
