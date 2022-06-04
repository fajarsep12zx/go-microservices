package main

import (
	"github.com/joho/godotenv"
	"github.com/micro/go-micro/v2/util/log"

	core "zebrax.id/product/dmaa/core/proto"
	health "zebrax.id/product/dmaa/core/proto/health"
	"zebrax.id/product/dmaa/core/server"
	"zebrax.id/product/dmaa/core/utils"
	"zebrax.id/product/dmaa/services/user/handler"
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
