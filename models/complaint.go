package models

type Complaint struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Email       string `gorm:"not null" json:"email"`
	Subject     string `gorm:"not null" json:"subject"`
	Description string `json:"description"`
	Status      string `gorm:"not null" json:"status"`
}
