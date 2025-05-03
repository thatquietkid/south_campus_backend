package models

import "gorm.io/gorm"

type CafeteriaItem struct {
    gorm.Model
    ItemName string  `json:"item_name"`
    Price    float64 `json:"price"`
}
