package testable

type TestCase struct {
	description string
	function    func(c TestContext)
}

func NewTestCase(description string, function func(c TestContext)) *TestCase {
	return &TestCase{
		description: description,
		function:    function,
	}
}

func (this *TestCase) Execute(c TestContext) {
	this.function(c)
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
