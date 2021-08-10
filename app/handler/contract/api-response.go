package contract

type StatusCode int

type APIResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

func CreateSuccessApiResponse(data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Data:    data,
	}
}

func CreateErrorApiResponse(err error) APIResponse {
	return APIResponse{
		Success: false,
		Error:   err.Error(),
		Data:    nil,
	}
}
