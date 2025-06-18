package models

type User struct {
	Email string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
