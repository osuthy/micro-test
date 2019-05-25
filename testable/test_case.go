package testable

type TestCase struct {
	function func()
}

func NewTestCase(function func()) *TestCase {
	return &TestCase{function: function}
}

func (this *TestCase) Execute() {
	this.function()
}

func (this *TestCase) HasNextTest() bool {
	return false
}

func (this *TestCase) NextTest() Testable {
	return nil
}
