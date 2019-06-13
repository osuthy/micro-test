package http

type ExpectedResponseBuilder struct {
	Status int
	Body string
}

func Status(statusCode int) ExpectedResponseBuilder {
	return ExpectedResponseBuilder{Status: statusCode}
}

func (this ExpectedResponseBuilder) TextPlain(body string) *Response {
	this.Body = body
	return this.Build()
}

func (this ExpectedResponseBuilder) Build() *Response {
	return NewResponse(this.Status, this.Body)
}

