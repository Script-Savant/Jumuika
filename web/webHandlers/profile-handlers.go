package webHandlers

import (
	"Jumuika/common/models"
	"Jumuika/common/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WebProfileController struct {
	DB *gorm.DB
}

func NewWebProfileController(db *gorm.DB) *WebProfileController {
	return &WebProfileController{db}
}

func (pc *WebProfileController) ViewProfile(c *gin.Context) {
	user := utils.GetCurrentUser(c, pc.DB)

	var profile models.Profile
	if err := pc.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "profile-view", gin.H{"error": "Error loading your profile"})
		return
	}

	var userReservedMeetings []models.RSVP
	if err := pc.DB.Where("user_id = ?", user.ID).Order("created_at desc").Find(&userReservedMeetings).Error; err != nil {
		c.HTML(http.StatusNotFound, "profile-view", gin.H{"error": "Error loading your reserved meetings"})
		return
	}

	var userCreatedMeetings []models.Meeting
	if err := pc.DB.Where("creator_id = ?", user.ID).Order("created_at desc").Find(&userCreatedMeetings).Error; err != nil {
		c.HTML(http.StatusNotFound, "profile-view", gin.H{"error": "Error loading the meetings you have created"})
		return
	}

	data := gin.H{
		"user":     user,
		"profile":  profile,
		"reserved": userReservedMeetings,
		"meetings": userCreatedMeetings,
	}

	c.HTML(http.StatusOK, "profile-view", data)
}

func (pc *WebProfileController) EditProfile(c *gin.Context) {
	user := utils.GetCurrentUser(c, pc.DB)

	var profile models.Profile
	if err := pc.DB.Where("user_id = ?", user.ID).First(&profile).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "profile-view", gin.H{"error": "Error loading your profile"})
		return
	}

	data := gin.H{
		"user":    user,
		"profile": profile,
	}

	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "profile-edit", data)
	}

	
}
