package subscription

// ErrorResponse represents an error returned by the API.
type ErrorResponse struct {
	Error string `json:"error" example:"something went wrong"`
}
