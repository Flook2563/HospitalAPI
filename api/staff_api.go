package api

import (
	"net/http"
	"os"
	"time"

	"github.com/Flook2563/Hospitalapi/config"
	"github.com/Flook2563/Hospitalapi/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT(staffID uint, hospitalID uint) (string, error) {
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"staff_id":    staffID,
		"hospital_id": hospitalID,
		"exp":         time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateStaff(c *gin.Context) {
	var InputStaff struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Hospital string `json:"hospital"`
	}

	if err := c.ShouldBindJSON(&InputStaff); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingStaff models.Staff
	if err := config.DB.Where("username = ?", InputStaff.Username).First(&existingStaff).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already taken"})
		return
	}

	var hospital models.Hospital
	if err := config.DB.Where("name = ?", InputStaff.Hospital).First(&hospital).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hospital Not Found !"})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(InputStaff.Password), 10)
	staff := models.Staff{
		Username:   InputStaff.Username,
		Password:   string(hashPassword),
		HospitalID: hospital.ID,
	}

	if err := config.DB.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//ดึงข้อมูลของโรงพยาบาลใส่ลงในตอน response
	if err := config.DB.Preload("Hospital").First(&staff, staff.ID).Error; err != nil {
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
		Hospital string `json:"hospital"`
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

	var hospital models.Hospital
	if err := config.DB.Where("id = ? AND name = ?", staff.HospitalID, InputStaff.Hospital).First(&hospital).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Hospital invalid !"})
		return
	}

	token, err := GenerateJWT(staff.ID, hospital.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token can not Create !"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login Success",
		"token":   token,
	})

}
