package microtest

import (
	"runtime"
	. "github.com/osuthy/micro-test/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()

var Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.AddSetUpFunction(setUpFunction)
}

func After(tearDownFunction func()) {
	testBuilder.AddTearDownFunction(tearDownFunction)
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
	return runtime.FuncForPC(pc).Name() == "github.com/osuthy/micro-test.Describe"
}

func It(params ...interface{}) interface{} {
	if description, ok := params[0].(string); ok {
		function, _ := params[1].(func(c TestContext))
		testBuilder.AddTestCase(description, function)
	} else {
		function, _ := params[0].(func(c TestContext))
		testBuilder.AddTestCase("", function)
	}
	return nil
}
