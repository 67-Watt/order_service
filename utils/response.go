package utils

// Response is the general response format for the API.
type Response struct {
	Status  string      `json:"status"`          // "success" or "error"
	Message string      `json:"message"`         // Descriptive message
	Data    interface{} `json:"data,omitempty"`  // Actual data payload (if any)
	Error   interface{} `json:"error,omitempty"` // Error details (if any)
}

// SuccessResponse creates a success response with data.
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  "success",
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

// ErrorResponse creates an error response with error details.
func ErrorResponse(message string, err interface{}) Response {
	return Response{
		Status:  "error",
		Message: message,
		Data:    nil,
		Error:   err,
	}
}
