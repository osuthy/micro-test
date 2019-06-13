package http

type Response struct {
	Status int
	Body string
}

func NewResponse(status int, body string) *Response {
	return &Response{Status: status, Body: body}
}
