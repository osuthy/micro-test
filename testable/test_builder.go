package testable

type TestBuilder struct {
	suites []*TestSuiteStruct
}

type TestSuiteStruct struct {
	setUpFunction *SetUpFunction
	testCases []*TestCase
}

func NewTestBuilder() *TestBuilder {
	return &TestBuilder{suites: []*TestSuiteStruct{}}
}

func (this *TestBuilder) AddTestSuite() {
	suite := &TestSuiteStruct{
		setUpFunction: nil,
		testCases: []*TestCase{},
	}
	this.suites = append([]*TestSuiteStruct{suite}, this.suites...)
}

func (this *TestBuilder) AddSetUpFunction(setUpFunction func()) {
	this.suites[0].setUpFunction = NewSetUpFunction(setUpFunction)
}

func (this *TestBuilder) AddTestCase(testFunction func()) {
	this.suites[0].testCases = append(this.suites[0].testCases, NewTestCase(testFunction))
}

func (this *TestBuilder) Build() *TestSuite {
	root := NewTestSuite(toTestable(this.suites[0].testCases), this.suites[0].setUpFunction)
	for i := 1; i < len(this.suites); i++ {
		testables := toTestable(this.suites[i].testCases)
		testables = append(testables, root)
		suite := NewTestSuite(testables, this.suites[i].setUpFunction)
		root = suite
	}
	return root
}

func toTestable(testCases []*TestCase) []Testable {
	testables := []Testable{}
	for _, testCase := range(testCases) {
		testables = append(testables, testCase)
	}
	return testables
}
