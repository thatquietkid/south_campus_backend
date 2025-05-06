package models

// Event represents the event model in the database.
type Event struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	EventDate   string `gorm:"not null" json:"event_date"`
	EventTime   string `json:"event_time"`
	Venue       string `json:"venue"`
	Description string `json:"description"`
	ContactEmail string `json:"contact_email"`
	Category    string `json:"category"`
}
