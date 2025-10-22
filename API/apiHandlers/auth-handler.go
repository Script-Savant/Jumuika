package apiHandlers

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

type ApiAuthController struct {
	DB *gorm.DB
}

type registerInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

type loginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewApiAuthController(db *gorm.DB) *ApiAuthController {
	return &ApiAuthController{db}
}

func (ac *ApiAuthController) Register(c *gin.Context) {
	var input registerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.Password != input.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	if !utils.ValidatePassword(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password does not meet the required threshold"})
		return
	}

	username := strings.TrimSpace(input.Username)

	var existingRecord models.User
	if err := ac.DB.Where("username = ?", username).First(&existingRecord).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A user with that username already exists"})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newRecord := models.User{
		Username:     username,
		PasswordHash: string(hashedPass),
	}

	if err := ac.DB.Create(&newRecord).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(http.StatusFound, "/login")
}

func (ac *ApiAuthController) Login(c *gin.Context) {
	var input loginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	username := strings.TrimSpace(input.Username)

	var user models.User
	if err := ac.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incorrect credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Incorrect credentials"})
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
