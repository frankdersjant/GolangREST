package Domain

type AppError struct {
	Code    int    `json:"omitempty"`
	Message string `json:"message"`
}
