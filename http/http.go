package http

import (
	"github.com/ShoichiroKitano/micro_test/runner"
)

func (this Response) AndResponseShouldBe(expected *Response) (int, string) {
	if !this.Equal(expected) {
		runner.TestRunner.Result = ""
		return 500, ""
	}
	runner.TestRunner.Result = "success"
	return 200, "test success"
}
