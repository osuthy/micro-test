package micro_test

import (
	. "github.com/ShoichiroKitano/micro_test/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()
var buildLock interface{} = nil

var	Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.BuildSetUp(setUpFunction)
}

func Feature(description string, testFunction func()) interface{} {
	lock := new(interface{})
	if(buildLock == nil) { buildLock = lock }

	testBuilder.BuildTestSuite()
	testFunction()

	if(buildLock == lock) {
		Suites = append(Suites, testBuilder.Pop())
		testBuilder = NewTestBuilder()
		buildLock = nil
	}
	return nil
}

func Test(description string, testFunction func()) interface{} {
	testBuilder.BuildTestCase(testFunction)
	return nil
}
