package response

import "time"

type CatResponse struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Race        string    `json:"race"`
	Sex         string    `json:"sex"`
	AgeInMonth  string    `json:"ageInMonth"`
	ImageUrls   []string  `json:"imageUrls"`
	Description string    `json:"description"`
	HasMatched  bool      `json:"hasMatched"`
	CreatedAt   time.Time `json:"createdAt"`
}
