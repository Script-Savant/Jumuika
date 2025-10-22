package apiRoutes

import (
	"Jumuika/API/apiHandlers"
	"Jumuika/common/config"
	"Jumuika/common/middleware"

	"github.com/gin-gonic/gin"
)

func SetupAuthApiRoutes(router *gin.Engine) {
	db := config.GetDB()
	controller := apiHandlers.NewApiAuthController(db)

	r := router.Group("/api") 
	{
		r.POST("/register", controller.Register)
		r.POST("/login", controller.Login)

		r.GET("/logout", middleware.AuthRequired(), apiHandlers.Logout)
	}
}