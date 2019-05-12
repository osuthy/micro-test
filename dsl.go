package micro_test

import (
	. "github.com/ShoichiroKitano/micro_test/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()
var buildLock interface{} = nil

var	Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.AddSetUpFunction(setUpFunction)
}

func Feature(description string, testFunction func()) interface{} {
	lock := new(interface{})
	if(buildLock == nil) { buildLock = lock }

	testBuilder.AddTestSuite()
	testFunction()

	if(buildLock == lock) {
		Suites = append(Suites, testBuilder.Build())
		testBuilder = NewTestBuilder()
		buildLock = nil
	}
	return nil
}

func Test(description string, testFunction func()) interface{} {
	testBuilder.AddTestCase(testFunction)
	return nil
}
