package dto

import "time"

type EmployeeData struct {
	ID       string
	Name     string
	Gender   string
	Birthday time.Time
	Salary   uint
}
