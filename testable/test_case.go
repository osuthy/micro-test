package testable

type TestCase struct {
	description string
	function func()
}

func NewTestCase(description string, function func()) *TestCase {
	return &TestCase{
		description: description,
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

func (this *TestCase) Descriptions() []string {
	return []string{this.description}
}
