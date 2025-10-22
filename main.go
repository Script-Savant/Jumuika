package main

import (
	"Jumuika/API/apiRoutes"
	"Jumuika/common/config"
	"Jumuika/ui/templates"
	"Jumuika/web/routes"
	"log"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// connect database
	config.DatabaseConnection()

	// router
	router := gin.Default()

	// load session and secret
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load session and secret", err)
	}

	// sessions
	store := cookie.NewStore([]byte(os.Getenv("SECRET")))
	router.Use(sessions.Sessions(os.Getenv("SESSION"), store))

	// templates
	router.HTMLRender = templates.SetupTemplates()

	// static files
	router.Static("/static", "./ui/static")

	// setup ui routes
	routes.SetupAuthRoutes(router)

	// setup api routes
	apiRoutes.SetupAuthApiRoutes(router)

	// start server
	log.Println("Server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}
