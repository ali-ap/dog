package internal

type FrameworkErrorCode int16

const (
	_ FrameworkErrorCode = iota
	PARAMETER_VALIDATION_ERROR
	GENERIC_ERROR
)

// swagger:model CommonError
type CommonError struct {
	// Status of the error
	// in: int16
	Status FrameworkErrorCode `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
}

func NewCommonError(status FrameworkErrorCode, message string) CommonError {
	return CommonError{
		status,
		message,
	}
}
