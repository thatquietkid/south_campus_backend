// models/bus_route.go
package models

type BusRoute struct {
	Number      string          `gorm:"primaryKey" json:"number"`
	Color       string          `json:"color"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Schedule    []ScheduleEntry `gorm:"type:jsonb" json:"schedule"`
}

// Set correct table name
func (BusRoute) TableName() string {
	return "bus_routes_combined"
}

type ScheduleEntry struct {
	Stop string `json:"stop"`
	Time string `json:"time"`
}
