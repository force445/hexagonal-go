package main

import (
	"hexagonal_v2/config"
	"hexagonal_v2/internal/adapter/repository"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.NewConfig()

	postgresClient, err := gorm.Open(postgres.Open(conf.BuildPostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	postgresRepository := repository.NewDB(postgresClient)
	_ = postgresRepository
}
