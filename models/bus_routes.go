package models

type Stop struct {
	RouteNumber string `json:"-"`
	StopName    string `json:"stop"`
	ArrivalTime string `json:"time"`
}

type Route struct {
	Number      string `gorm:"primaryKey" json:"number"`
	Color       string `json:"color"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    []Stop `gorm:"-" json:"schedule"` // manually append in handler
}