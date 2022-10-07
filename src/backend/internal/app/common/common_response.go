package internal

// swagger:model CommonResponse
type CommonResponse struct {

	// Message of the response
	// in: string
	Message string `json:"message"`
}

func NewCommonResponse(message string) CommonResponse {
	return CommonResponse{
		Message: message,
	}
}
