package models

type BusRoute struct {
	Number      string          `gorm:"primaryKey" json:"number"`
	Color       string          `json:"color"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedule    []ScheduleEntry `gorm:"type:jsonb" json:"schedule"`
}

// ScheduleEntry represents each item in the schedule array
type ScheduleEntry struct {
	Stop string `json:"stop"`
	Time string `json:"time"`
}
