package utils

import (
	"Jumuika/common/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCurrentUser(c *gin.Context, db *gorm.DB) *models.User {
	session := sessions.Default(c)
	userID := session.Get("user_id")

	if userID != nil {
		var user models.User
		db.First(&user, userID)
		return &user
	}
	c.Redirect(http.StatusFound, "/login")
	return nil
}