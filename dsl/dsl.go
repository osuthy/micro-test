package dsl

import (
	"github.com/ShoichiroKitano/micro_test/runner"
)

func Test(description string, test func()) interface{} {
	runner.TestMethod = test
	return nil
}
