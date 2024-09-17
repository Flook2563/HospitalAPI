package api

import (
	"net/http"

	"github.com/Flook2563/Hospitalapi/config"
	"github.com/Flook2563/Hospitalapi/models"
	"github.com/gin-gonic/gin"
)

func SearchPatient(c *gin.Context) {
	hospitalID, _ := c.Get("hospital_id")

	var patients []models.Patient
	query := config.DB.Where("hospital_id = ?", hospitalID)

	if nationalID := c.Query("national_id"); nationalID != "" {
		query = query.Where("national_id = ?", nationalID)
	}

	if passportID := c.Query("passport_id"); passportID != "" {
		query = query.Where("passport_id = ?", passportID)
	}

	if firstName := c.Query("first_name"); firstName != "" {
		query = query.Where("first_name_th LIKE ? OR first_name_en LIKE ?", "%"+firstName+"%", "%"+firstName+"%")
	}

	if middleName := c.Query("middle_name"); middleName != "" {
		query = query.Where("middle_name_th LIKE ? OR middle_name_en LIKE ?", "%"+middleName+"%", "%"+middleName+"%")
	}

	if lastName := c.Query("last_name"); lastName != "" {
		query = query.Where("last_name_th LIKE ? OR last_name_en LIKE ?", "%"+lastName+"%", "%"+lastName+"%")
	}

	if dateOfBirth := c.Query("date_of_birth"); dateOfBirth != "" {
		query = query.Where("date_of_birth = ?", dateOfBirth)
	}

	if phoneNumber := c.Query("phone_number"); phoneNumber != "" {
		query = query.Where("phone_number = ?", phoneNumber)
	}

	if email := c.Query("email"); email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	if err := query.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"patients": patients})
}
