package http

func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		this.differences.Push("assert is fail!!!!!!!!!!!1")
	}
}
