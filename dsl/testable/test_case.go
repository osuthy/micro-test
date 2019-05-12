package testable

type TestCase struct {
	canExecute bool
	function func()
}

func NewTestCase(function func()) *TestCase {
	return &TestCase{canExecute: true, function: function}
}

func (this *TestCase) Execute() {
	this.function()
	this.canExecute = false
}

func (this *TestCase) HasUnexecutedTest() bool {
	return this.canExecute
}
