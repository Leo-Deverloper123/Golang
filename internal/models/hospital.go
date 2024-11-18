package models

import (
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name     string `json:"name"`
	APIUrl   string `json:"api_url"`
	APIToken string `json:"api_token"`
}