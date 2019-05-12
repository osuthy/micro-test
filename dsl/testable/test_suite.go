package testable

type TestSuite struct {
	tests []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite() *TestSuite {
	return new(TestSuite)
}

func NewTestSuite2(setUpFunction *SetUpFunction, tests []Testable) *TestSuite {
	return &TestSuite{tests: tests, setUpFunction: setUpFunction}
}

func (this *TestSuite) Execute() {
	for _, test := range this.tests {
		if test.HasUnexecutedTest() {
			if this.setUpFunction != nil {
				this.setUpFunction.Execute()
			}
			test.Execute()
			return
		}
	}
}

func (this *TestSuite) Add(test Testable) {
	this.tests = append(this.tests, test)
}

func (this *TestSuite) HasUnexecutedTest() bool {
	for _, test := range this.tests {
		if test.HasUnexecutedTest() {
			return true
		}
	}
	return false
}
