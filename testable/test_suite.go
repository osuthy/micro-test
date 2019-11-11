package testable

type TestSuite struct {
	description          string
	tests         []Testable
	setUpFunction *SetUpFunction
	tearDownFunction *TearDownFunction
}

func NewTestSuite(description string, tests []Testable, setUpFunction *SetUpFunction, tearDownFunction *TearDownFunction) *TestSuite {
	return &TestSuite{
		description:         description,
		tests:         tests,
		setUpFunction: setUpFunction,
		tearDownFunction: tearDownFunction,
	}
}

func (this *TestSuite) Execute(c TestContext) {
	if this.setUpFunction != nil {
		this.setUpFunction.Execute()
	}
	this.tests[0].Execute(c)
	if this.tearDownFunction != nil {
		this.tearDownFunction.Execute()
	}
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
		return NewTestSuite(this.description, tests, this.setUpFunction, this.tearDownFunction)
	}
	tests = append(tests, this.tests[1:]...)
	return NewTestSuite(this.description, tests, this.setUpFunction, this.tearDownFunction)
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

