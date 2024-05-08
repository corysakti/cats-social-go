package relation

import "time"

type UserCatMatchRelation struct {
	Id          int32
	Email       string
	Password    string
	UserId      string
	CatId       string
	CatName     string
	UserName    string
	Race        string
	Sex         string
	Description string
	AgeInMonth  int32
	ImageUrls   string
	CreatedAt   time.Time
}
