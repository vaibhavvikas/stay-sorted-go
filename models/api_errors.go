package models

type AppError struct {
	Code    int    `json:"code"`
	Param   string `json:"param,omitempty"`
	Message string `json:"message"`
}

func (m *AppError) Error() string {
	return m.Message
}
