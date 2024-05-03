package response

type ResponseTemplate struct {
	Code   int
	Status string
	Data   interface{}
}
