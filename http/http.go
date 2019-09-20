package http

import (
	"github.com/ShoichiroKitano/micro_test/runner"
)

func (this Response) AndResponseShouldBe(expected *Response) {
	if !this.Equal(expected) {
		runner.Queue.Push("assert is fail!!!!!!!!!!!1")
		runner.TestRunner.Result = ""
	} else {
		runner.TestRunner.Result = "success"
	}
}
