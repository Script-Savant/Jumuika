package routes

import (
	"Jumuika/common/config"
	"Jumuika/common/middleware"
	"Jumuika/web/webHandlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	db := config.GetDB()
	authController := webHandlers.NewAuthController(db)

	router.GET("/register", authController.Register)
	router.POST("/register", authController.Register)
	router.GET("/login", authController.Login)
	router.POST("/login", authController.Login)

	r := router.Group("/")
	r.Use(middleware.AuthRequired())
	{
		r.GET("/logout", webHandlers.Logout)
	}

}