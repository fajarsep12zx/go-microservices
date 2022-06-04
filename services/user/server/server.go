package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2/util/log"

	core "go-microservices/core/proto"
	health "go-microservices/core/proto/health"
	"go-microservices/core/server"
	"go-microservices/core/utils"
	"go-microservices/services/user/handler"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found ..")
	}

	log.Info("Data Source : " + utils.GetDatasourceInfo())
}

func main() {
	service := server.NewGRPCServer(server.UserService)
	service.Init()

	srv := service.Server()
	core.RegisterUserHandler(srv, new(handler.Handler))
	health.RegisterHealthHandler(srv, server.NewHealthCheck())

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
