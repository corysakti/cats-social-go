package request

type CategoryCreateRequest struct {
	Name string `validate:"required,min=3,max=100"`
}
