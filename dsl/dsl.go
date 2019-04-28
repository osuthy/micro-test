package dsl

type Testable interface {
	Execute()
}

var	Suites []Testable = []Testable{}

type TestCase struct {
	function func()
}

func NewTestCase(function func()) *TestCase {
	testCase := new(TestCase)
	testCase.function = function
	return testCase
}

func (this *TestCase) Execute() {
	this.function()
}

type SetUpFunction struct {
	function func()
}

func NewSetUpFunction(function func()) *SetUpFunction {
	setUpFunction := new(SetUpFunction)
	setUpFunction.function = function
	return setUpFunction
}

func (this *SetUpFunction) Execute() {
	this.function()
}

type TestSuite struct {
	tests []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite() *TestSuite {
	return new(TestSuite)
}

func (this *TestSuite) Execute() {
	for _, test := range this.tests {
		if this.setUpFunction != nil {
			this.setUpFunction.Execute()
		}
		test.Execute()
	}
}

func (this *TestSuite) Add(test Testable) {
	this.tests = append(this.tests, test)
}

type testSuiteBuilder struct {
	suites []*TestSuite
}

func NewTestSuiteBuilder() *testSuiteBuilder {
	builder := new(testSuiteBuilder)
	builder.suites = []*TestSuite{}
	return builder
}

var TestSuiteBuilder *testSuiteBuilder = NewTestSuiteBuilder()

func (this *testSuiteBuilder) BuildTestSuite() {
	this.suites = append([]*TestSuite{NewTestSuite()}, this.suites...)
}

func (this *testSuiteBuilder) BuildSetUp(setUpFunction func()) {
	this.suites[0].setUpFunction = NewSetUpFunction(setUpFunction)
}

func (this *testSuiteBuilder) BuildTestCase(testFunction func()) {
	this.suites[0].Add(NewTestCase(testFunction))
}

func (this *testSuiteBuilder) Pop() Testable {
	last := this.suites[0]
	for i := 1; i < len(this.suites); i++ {
		this.suites[i].Add(last)
		last = this.suites[i]
	}
	return last
}

func Before(setUpFunction func()) {
	TestSuiteBuilder.BuildSetUp(setUpFunction)
}

func Feature(description string, testFunction func()) interface{} {
	TestSuiteBuilder.BuildTestSuite()
	testFunction()
	Suites = append(Suites, TestSuiteBuilder.Pop())
	TestSuiteBuilder = NewTestSuiteBuilder()
	return nil
}

func Test(description string, testFunction func()) interface{} {
	TestSuiteBuilder.BuildTestCase(testFunction)
	return nil
}
