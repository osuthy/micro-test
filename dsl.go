package micro_test

import (
	"runtime"
	. "github.com/ShoichiroKitano/micro_test/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()

var Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.AddSetUpFunction(setUpFunction)
}

func Describe(description string, testFunction func()) interface{} {
	testBuilder.AddTestSuite(description)
	testFunction()
	if !calledByDescribeForDescribe() {
		Suites = append(Suites, testBuilder.Build())
		testBuilder = NewTestBuilder()
	}
	return nil
}

func calledByDescribeForDescribe() bool {
	pc, _, _, _ := runtime.Caller(3)
	return runtime.FuncForPC(pc).Name() == "github.com/ShoichiroKitano/micro_test.Describe"
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
