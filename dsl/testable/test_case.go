package testable

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
