package testable

type TestSuite struct {
	tests []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite() *TestSuite {
	return new(TestSuite)
}

func NewTestSuite2(tests []Testable, setUpFunction *SetUpFunction) *TestSuite {
	return &TestSuite {
		tests: tests,
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
	if(test != nil) { tests = append(tests, test) }
	if(len(this.tests) == 1) {
		if(len(tests) == 0) {
			return nil
		}
		return NewTestSuite2(tests, this.setUpFunction)
	}
	tests = append(tests, this.tests[1:]...)
	return NewTestSuite2(tests, this.setUpFunction)
}

func (this *TestSuite) Add(test Testable) {
	this.tests = append(this.tests, test)
}

func (this *TestSuite) HasNextTest() bool {
	if len(this.tests) == 1 {
		return this.tests[0].HasNextTest()
	}
	return true
}

