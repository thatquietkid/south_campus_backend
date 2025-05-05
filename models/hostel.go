package models

// Hostel represents the hostel model in the database.
type Hostel struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `gorm:"not null" json:"name"`
	Type          string `gorm:"not null" json:"type"`
	Address       string `gorm:"type:text" json:"address"`
	Warden        string `json:"warden"`
	WardenContact string `json:"warden_contact"`
	Caretaker     string `json:"caretaker"`
	CaretakerContact string `json:"caretaker_contact"`
	Capacity      int    `json:"capacity"`
	Available     int    `json:"available"`
	Facilities    string `gorm:"type:text" json:"facilities"`
}