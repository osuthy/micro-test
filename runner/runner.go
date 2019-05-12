package runner

import (
	"github.com/ShoichiroKitano/micro_test/dsl"
)

type testRunner struct {
	Result string
}

var TestRunner testRunner = testRunner{}

func Run() {
	for _, suite := range dsl.Suites {
		for suite.HasUnexecutedTest() {
			suite.Execute()
		}
	}
}
