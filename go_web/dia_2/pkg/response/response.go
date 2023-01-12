package response

import "errors"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExist        = errors.New("already exists")
	ErrInternalServerError = errors.New("internal server error")
	ErrInvalidId           = errors.New("invalid id")
	ErrUnauthorized        = errors.New("unauthorized")
)

func Ok(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

func Err(err error) *Response {
	return &Response{Message: err.Error(), Data: nil}
}
