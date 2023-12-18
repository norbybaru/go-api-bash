package main

import (
	"dancing-pony/cmd/server"
	_ "github.com/joho/godotenv/autoload"
)

// @Dancing Pony API
// @version 1.0
// @description This is a API documentation how how to interact with this service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email norbybaru@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	server := server.NewApp()

	server.Start()
}
