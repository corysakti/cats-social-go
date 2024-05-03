package request

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
