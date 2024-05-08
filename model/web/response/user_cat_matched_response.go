package response

import "time"

type UserCatMatchedResponse struct {
	Id             int          `json:"id"`
	IssuedBy       UserResponse `json:"issuedBy"`
	MatchCatDetail CatResponse  `json:"matchCatDetail"`
	UserCatDetail  CatResponse  `json:"userCatDetail"`
	Message        string       `json:"message"`
	CreatedAt      time.Time    `json:"createdAt"`
}
