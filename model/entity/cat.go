package entity

import "time"

type Cat struct {
	Id          int32
	Name        string
	Race        string
	Sex         string
	Description string
	AgeInMonth  int32
	ImageUrls   string
	CreatedAt   time.Time
}
