package request

type CatMatch struct {
	MatchCatID string `json:"matchCatId" validate:"required"`
	UserCatID  string `json:"userCatId" validate:"required"`
	Message    string `json:"message" validate:"required,min=5,max=120"`
}
