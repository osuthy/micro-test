package http

import (
	"github.com/osuthy/micro-test/runner"
)

func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		runner.Diffs.Push("assert is fail!!!!!!!!!!!1")
	}
}
