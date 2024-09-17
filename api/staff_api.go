package api

import (
	"net/http"

	"github.com/Flook2563/Hospitalapi/config"
	"github.com/Flook2563/Hospitalapi/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateStaff(c *gin.Context) {
	var staff models.Staff
	if err := c.ShouldBindJSON(&staff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(staff.Password), 10)
	staff.Password = string(hashPassword)

	if err := config.DB.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"staff": staff})

}

func LoginStaff(c *gin.Context) {
	var staff models.Staff
	var InputStaff struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&InputStaff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Where("username = ?", InputStaff.Username).First(&staff).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username invalid !"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(InputStaff.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password invalid !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login Success"})

}
