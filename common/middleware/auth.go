package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			session.Set("redirect_url", c.Request.URL.String())
			session.Save()

			if isAPIRequest(c) {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized: Please log in first",
				})
			} else {
				c.Redirect(http.StatusFound, "/login")
			}

			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		roleVal := session.Get("user_role")

		role, ok := roleVal.(string)
		if !ok || role != "admin" {
			if isAPIRequest(c) {
				c.JSON(http.StatusForbidden, gin.H{
					"error": "Forbidden: Admin access required",
				})
			} else {
				c.HTML(http.StatusForbidden, "error", gin.H{
					"error": "Unauthorized access",
				})
			}

			c.Abort()
			return
		}

		c.Next()
	}
}

func isAPIRequest(c *gin.Context) bool {
	accept := c.GetHeader("Accept")
	authHeader := c.GetHeader("Authorization")

	return strings.Contains(accept, "application/json") || authHeader != ""
}
