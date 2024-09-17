package models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Username   string   `gorm:"unique;not null" json:"username"` // ชื่อผู้ใช้
	Password   string   `gorm:"not null" json:"password"`
	HospitalID uint     `json:"hospital_id"`
	Hospital   Hospital `json:"hospital"`
}

type Hospital struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
}
