package testable

type TestCase struct {
	name string
	function func()
}

func NewTestCase(name string, function func()) *TestCase {
	return &TestCase{
		name: name,
		function: function,
	}
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
