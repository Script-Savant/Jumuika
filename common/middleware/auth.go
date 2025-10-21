package middleware

import (
	"net/http"

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

			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role := session.Get("user_role").(string)

		if role != "admin" {
			c.HTML(http.StatusForbidden, "error", gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
