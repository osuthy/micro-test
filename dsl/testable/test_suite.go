package testable

type TestSuite struct {
	tests []Testable
	setUpFunction *SetUpFunction
}

func NewTestSuite() *TestSuite {
	return new(TestSuite)
}

func (this *TestSuite) Execute() {
	for _, test := range this.tests {
		if this.setUpFunction != nil {
			this.setUpFunction.Execute()
		}
		test.Execute()
	}
}

func (this *TestSuite) Add(test Testable) {
	this.tests = append(this.tests, test)
}
