package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/octadsp/server-home-test/database"
	connectiondb "github.com/octadsp/server-home-test/pkg/connectionDB"
	"github.com/octadsp/server-home-test/routes"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	connectiondb.DatabaseInit()
	database.RunMigration()

	router := gin.New()
	routes.RouteInit(router.Group("/api/v1"))

	// var PORT = os.Getenv("PORT")

	fmt.Println("Server running localhost :" + "5000")
	// g.Logger.Fatal(g.Start(":" + PORT))
	router.Use(gin.Logger())

	err := router.Run(":" + "5000")
	if err != nil {
		log.Fatal(err)
	}
}
