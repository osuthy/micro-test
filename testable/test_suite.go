package testable

type TestSuite struct {
	name          string
	tests         []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite(name string, tests []Testable, setUpFunction *SetUpFunction) *TestSuite {
	return &TestSuite{
		name:         name,
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
		return NewTestSuite("", tests, this.setUpFunction)
	}
	tests = append(tests, this.tests[1:]...)
	return NewTestSuite("", tests, this.setUpFunction)
}

func (this *TestSuite) HasNextTest() bool {
	if len(this.tests) == 1 {
		return this.tests[0].HasNextTest()
	}
	return true
}
