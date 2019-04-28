package dsl

import (
	"github.com/ShoichiroKitano/micro_test/runner"
)

func Test(description string, testFunction func()) interface{} {
	runner.TestFunctions = append(runner.TestFunctions, testFunction)
	return nil
}
