package http

type APIResponse struct {
	Success bool        `json:"success"`
	Errors  []string    `json:"errors"`
	Data    interface{} `json:"data"`
}

func CreateSuccessApiResponse(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Errors:  nil,
		Data:    data,
	}
}