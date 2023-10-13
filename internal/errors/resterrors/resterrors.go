package resterrors

type RestError struct {
	StatusCode   int    `json:"status_code"`
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

// CreateNew creates new RestError object
func CreateNew(StatusCode int, ErrorCode string, ErrorMessage string) *RestError {
	return &RestError{
		StatusCode:   StatusCode,
		ErrorCode:    ErrorCode,
		ErrorMessage: ErrorMessage,
	}
}
