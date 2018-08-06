package respond

type MessagePayload struct{
	Message string `json:"message"`
}

type ErrorPayload struct{
	Error string `json:"error"`
}
