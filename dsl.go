package micro_test

import (
	. "github.com/ShoichiroKitano/micro_test/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()
var buildLock interface{} = nil

var Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.AddSetUpFunction(setUpFunction)
}

func Describe(description string, testFunction func()) interface{} {
	lock := new(interface{})
	if buildLock == nil {
		buildLock = lock
	}

	testBuilder.AddTestSuite(description)
	testFunction()

	if buildLock == lock {
		Suites = append(Suites, testBuilder.Build())
		testBuilder = NewTestBuilder()
		buildLock = nil
	}
	return nil
}

func It(params ...interface{}) interface{} {
	if description, ok := params[0].(string); ok {
		function, _ := params[1].(func())
		testBuilder.AddTestCase(description, function)
	} else {
		function, _ := params[0].(func())
		testBuilder.AddTestCase("", function)
	}
	return nil
}
