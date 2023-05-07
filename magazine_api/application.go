package main

import (
	"magazine_api/bootstrap"
	"magazine_api/lib"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

// @title        Magazine API Documentation
// @version      v1
// @description  This is an API documentation of Magazine API

// @contact.Name  21557870 - Suman Chhetri

// @host      localhost:5005
// @BasePath  /api/v1

// @securityDefinitions.apikey  BearerAuth
// @in                          header
// @name                        Authorization
func main() {
	godotenv.Load()

	logger := lib.GetLogger()
	fx.New(bootstrap.Module, fx.Logger(logger.GetFxLogger())).Run()
}
