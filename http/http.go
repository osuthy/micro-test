package http

import (
	"github.com/ShoichiroKitano/micro_test/runner"
)

func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		runner.Diffs.Push("assert is fail!!!!!!!!!!!1")
	}
}
