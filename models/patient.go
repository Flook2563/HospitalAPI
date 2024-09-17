package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	First_Name_Th  	string 	`json:"first_name_th"`
	Middle_Name_Th 	string 	`json:"middle_name_th"`
	Last_Name_Th   	string 	`json:"last_name_th"`
	First_Name_En  	string 	`json:"first_name_en"`
	Middle_Name_En 	string 	`json:"middle_name_en"`
	Last_Name_En   	string 	`json:"last_name_en"`
	Date_Of_Birth  	string 	`json:"date_of_birth"`
	Patient_HN     	string 	`json:"patient_hn"`
	National_ID    	string 	`json:"national_id"`
	Passport_ID    	string 	`json:"passport_id"`
	Phone_Number   	string 	`json:"phone_number"`
	Email          	string 	`json:"email"`
	Gender         	string 	`json:"gender"`
	HospitalID 		uint	`json:"hospital_id"`
}
