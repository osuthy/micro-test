package runner

import (
	"github.com/ShoichiroKitano/micro_test"
)

type testRunner struct {
	Result string
}

var TestRunner testRunner = testRunner{}

func Run() {
	for _, suite := range micro_test.Suites {
		s := suite
		for {
			s.Execute()
			if !s.HasNextTest() { break }
			s = s.NextTest()
		}
	}
}
