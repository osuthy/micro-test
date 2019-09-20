package testable

type TestSuite struct {
	description          string
	tests         []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite(description string, tests []Testable, setUpFunction *SetUpFunction) *TestSuite {
	return &TestSuite{
		description:         description,
		tests:         tests,
		setUpFunction: setUpFunction,
	}
}

func (this *TestSuite) Execute() {
	if this.setUpFunction != nil {
		this.setUpFunction.Execute()
	}
	this.tests[0].Execute()
}

func (this *TestSuite) NextTest() Testable {
	tests := []Testable{}
	test := this.tests[0].NextTest()
	if test != nil {
		tests = append(tests, test)
	}
	if len(this.tests) == 1 {
		if len(tests) == 0 {
			return nil
		}
		return NewTestSuite(this.description, tests, this.setUpFunction)
	}
	tests = append(tests, this.tests[1:]...)
	return NewTestSuite(this.description, tests, this.setUpFunction)
}

func (this *TestSuite) HasNextTest() bool {
	if len(this.tests) == 1 {
		return this.tests[0].HasNextTest()
	}
	return true
}

func (this *TestSuite) Descriptions() []string {
	return append([]string{this.description}, this.tests[0].Descriptions()...)
}

