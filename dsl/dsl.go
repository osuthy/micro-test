package dsl

import (
	. "github.com/ShoichiroKitano/micro_test/dsl/testable"
)

var testBuilder *TestBuilder = NewTestBuilder()

var	Suites []Testable = []Testable{}

func Before(setUpFunction func()) {
	testBuilder.BuildSetUp(setUpFunction)
}

func Feature(description string, testFunction func()) interface{} {
	testBuilder.BuildTestSuite()
	testFunction()
	Suites = append(Suites, testBuilder.Pop())
	testBuilder = NewTestBuilder()
	return nil
}

func Test(description string, testFunction func()) interface{} {
	testBuilder.BuildTestCase(testFunction)
	return nil
}
