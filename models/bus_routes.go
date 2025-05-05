package models

import "gorm.io/gorm"

type BusRoute struct {
	gorm.Model
	Number      string `gorm:"primaryKey"`
	Color       string
	Name        string
	Description string
	Schedule    []map[string]string `gorm:"type:jsonb"`
}