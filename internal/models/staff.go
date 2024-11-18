package models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	Username    string `json:"username" gorm:"uniqueIndex"`
	Password    string `json:"-"`
	HospitalID  uint   `json:"hospital_id"`
	Hospital    Hospital
}