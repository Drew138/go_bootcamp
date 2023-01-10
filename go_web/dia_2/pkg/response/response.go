package response

import "errors"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var (
	ErrNotFound = errors.New("not found")
)

func Ok(message string, data interface{}) *Response {
	return &Response{Message: message, Data: data}
}

func Err(err error) *Response {
	return &Response{Message: err.Error(), Data: nil}
}
