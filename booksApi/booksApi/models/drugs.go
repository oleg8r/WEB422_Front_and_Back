package models

import "gorm.io/gorm"

type Drug struct {
	gorm.Model
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
	Quantity    int    `form:"quantity" json:"quantity"`
}
