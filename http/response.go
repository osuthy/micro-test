package http

import (
	. "github.com/osuthy/micro-test"
)

type Response struct {
	Status int
	Body   string
	differences *Differences
}

func NewResponse(status int, body string) *Response {
	return &Response{Status: status, Body: body}
}

func (this Response) Equal(other *Response) bool {
	if this.Status != other.Status {
		return false
	}
	if this.Body != other.Body {
		return false
	}
	return true
}
