package response

type DataResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
