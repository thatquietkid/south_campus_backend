package models

type CourseAttendance struct {
	Title               string  `json:"title"`
	Code                string  `json:"code"`
	AttendancePercentage float64 `json:"attendance_percentage"`
}
