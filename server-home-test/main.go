package main

import (
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
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

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"X-Requested-With", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	routes.RouteInit(router.Group("/api/v1"))

	fmt.Println("Server running localhost :" + "5000")
	// g.Logger.Fatal(g.Start(":" + PORT))
	router.Use(gin.Logger())

	router.Static("/uploads", "./uploads")

	err := router.Run(":" + "5000")
	if err != nil {
		log.Fatal(err)
	}
}
