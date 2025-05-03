package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title string `json:"title"`
	Code  string `json:"code"`
}
