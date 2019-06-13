package http

import(
	"github.com/ShoichiroKitano/micro_test/runner"
)

func (this Response) AndResponseShouldBe(expected *Response) (int, string) {
	if this.IsSame(expected) {
		runner.TestRunner.Result = "success"
		return 200, "test success"
	} else {
		runner.TestRunner.Result = ""
		return 500, ""
	}
}

func (this Response) IsSame(expected *Response) bool {
	if this.Status == expected.Status && this.Body == expected.Body {
		return true
	} else {
		return false
	}
}
