package main

import (
	"hexagonal_v2/config"
	"hexagonal_v2/internal/adapter/handler"
	"hexagonal_v2/internal/adapter/repository"
	"hexagonal_v2/internal/core/service"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.NewConfig()
	port := conf.GetAppPort()

	postgresClient, err := gorm.Open(postgres.Open(conf.BuildPostgresDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(postgresClient)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	airpodRepo := repository.NewAirpodRepository(postgresClient)
	airpodService := service.NewAirpodService(airpodRepo, userRepo)
	airpodHandler := handler.NewAirpodHandler(airpodService)
	locationRepo := repository.NewLocationRepository(postgresClient)
	locationService := service.NewLocationService(locationRepo)
	locationHandler := handler.NewLocationHandler(locationService)

	app := fiber.New()

	v1 := app.Group("/api/v1")

	airpodservices := v1.Group("/airpods")
	airpodservices.Post("", airpodHandler.CreateAirpod)
	airpodservices.Get("", airpodHandler.GetAirpods)
	airpodservices.Get("/:id", airpodHandler.GetAirpodByID)
	airpodservices.Get("/user/:id", airpodHandler.GetAirpodByUserID)
	airpodservices.Put("/:id", airpodHandler.UpdateAirpod)

	locationservices := v1.Group("/locations")
	locationservices.Post("", locationHandler.CreateLocation)
	locationservices.Get("", locationHandler.GetLocations)
	locationservices.Get("/:id", locationHandler.GetLocationByID)
	locationservices.Put("/:id", locationHandler.UpdateLocation)
	locationservices.Delete("/:id", locationHandler.DeleteLocation)

	userservices := v1.Group("/users")
	userservices.Post("", userHandler.CreateUser)
	userservices.Get("", userHandler.GetUsers)
	userservices.Get("/:id", userHandler.GetUserByID)
	userservices.Put("/:id", userHandler.UpdateUser)
	userservices.Delete("/:id", userHandler.DeleteUser)

	log.Fatal(app.Listen(":" + port))

	log.Printf("Server started on port %s", port)
}
