package main

import (
	"fmt"
	"log"
	"os"

	// "super-feed-service/cmd/config"

	"super-feed-service/cmd/config"
	"super-feed-service/cmd/docs"
	"super-feed-service/cmd/routers"

	_ "super-feed-service/cmd/docs"

	"github.com/joho/godotenv"
)

// @title                      Feed API documentation
// @version                    1.0.0
// @BasePath                   /api/v1
// @securityDefinitions.apikey BearerAuth
// @in                         header
// @name                       Authorization
func main() {
	routersInit := routers.InitRouter()

	err := godotenv.Load(".env")
	if err != nil {

		log.Fatal("Error while reading envs")
	}

	// Getting dynamic URL for Swagger Doc
	docs.SwaggerInfo.Host = os.Getenv("URL")
	docs.SwaggerInfo.Schemes = []string{os.Getenv("SCHEME")}

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	config.Connect()
	routersInit.Run(port)
}
