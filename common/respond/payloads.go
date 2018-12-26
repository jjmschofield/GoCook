package respond

// swagger:response
type MessageResponse struct {
	Body MessagePayload
}

// swagger:model
type MessagePayload struct {
	Message string `json:"message"`
}

// swagger:response
type ErrorResponse struct {
	Body ErrorPayload
}

// swagger:model
type ErrorPayload struct {
	Error string `json:"error"`
}
