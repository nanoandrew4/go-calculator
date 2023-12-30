package cloud

type CalculateRequest struct {
	Operation string `json:"operation"`
}

type CalculateResponse struct {
	Result string `json:"result,omitempty"`
}

type CalculateError struct {
	Error string `json:"error,omitempty"`
}
