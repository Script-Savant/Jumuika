package webHandlers

import (
	"Jumuika/common/models"
	"Jumuika/common/utils"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db}
}

func (ac *AuthController) Register(c *gin.Context) {
	if c.Request.Method == "GET" {
		data := gin.H{
			"Title": "register",
		}
		c.HTML(http.StatusOK, "register", data)
	}

	username := strings.TrimSpace(c.PostForm("username"))
	// email := strings.TrimSpace(c.PostForm("email"))
	password := strings.TrimSpace(c.PostForm("password"))
	password2 := strings.TrimSpace(c.PostForm("password2"))

	if password != password2 {
		c.HTML(http.StatusBadRequest, "register", gin.H{"error": "Passwords do not match"})
		return
	}

	if !utils.ValidatePassword(password) {
		c.HTML(http.StatusBadRequest, "register", gin.H{"error": "Password failed to reach recommended threshold"})
		return
	}

	if err := ac.DB.Where("username = ?", username).Error; err == nil {
		c.HTML(http.StatusBadRequest, "register", gin.H{"error": "A user with that username already exists"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "register", gin.H{"error": "Failed to encrypt password"})
		return
	}

	tx := ac.DB.Begin()

	newUserRecord := models.User{
		Username:     username,
		PasswordHash: string(hashedPass),
	}

	if err := tx.Create(&newUserRecord).Error; err != nil {
		tx.Rollback()
		c.HTML(http.StatusInternalServerError, "register", gin.H{"error": "Failed to create new user"})
		return
	}

	newProfileRecord := models.Profile{
		UserID: newUserRecord.ID,
	}

	if err := tx.Create(&newProfileRecord).Error; err != nil {
		tx.Rollback()
		c.HTML(http.StatusInternalServerError, "register", gin.H{"error": "Failed to create user profile"})
		return
	}

	tx.Commit()

	c.Redirect(http.StatusFound, "/login")
}

func (ac *AuthController) Login(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "login", gin.H{"Title": "login"})
		return
	}

	username := strings.TrimSpace(c.PostForm("username"))
	password := strings.TrimSpace(c.PostForm("password"))

	var user models.User
	if err := ac.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.HTML(http.StatusNotFound, "login", gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		c.HTML(http.StatusNotFound, "login", gin.H{"error": "Invalid credentials"})
		return
	}

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	session.Set("user_role", user.Role)
	session.Save()

	c.Redirect(http.StatusFound, "/home")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(http.StatusFound, "/home")
}
