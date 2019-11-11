package testable

type TestBuilder struct {
	suites []*TestSuiteStruct
}

type TestSuiteStruct struct {
	description string
	setUpFunction *SetUpFunction
	tearDownFunction *TearDownFunction
	testCases     []*TestCase
}

func NewTestBuilder() *TestBuilder {
	return &TestBuilder{suites: []*TestSuiteStruct{}}
}

func (this *TestBuilder) AddTestSuite(description string) {
	suite := &TestSuiteStruct{
		description: description,
		setUpFunction: nil,
		testCases:     []*TestCase{},
	}
	this.suites = append([]*TestSuiteStruct{suite}, this.suites...)
}

func (this *TestBuilder) AddSetUpFunction(setUpFunction func()) {
	this.suites[0].setUpFunction = NewSetUpFunction(setUpFunction)
}

func (this *TestBuilder) AddTearDownFunction(tearDownFunction func()) {
	this.suites[0].tearDownFunction = NewTearDownFunction(tearDownFunction)
}

func (this *TestBuilder) AddTestCase(description string, testFunction func(c TestContext)) {
	this.suites[0].testCases = append(this.suites[0].testCases, NewTestCase(description, testFunction))
}

func (this *TestBuilder) Build() *TestSuite {
	root := NewTestSuite(
		this.suites[0].description,
		toTestable(this.suites[0].testCases),
		this.suites[0].setUpFunction,
		this.suites[0].tearDownFunction)
	for i := 1; i < len(this.suites); i++ {
		testables := toTestable(this.suites[i].testCases)
		testables = append(testables, root)
		suite := NewTestSuite(
			this.suites[i].description,
			testables,
			this.suites[i].setUpFunction,
			this.suites[i].tearDownFunction)
		root = suite
	}
	return root
}

func toTestable(testCases []*TestCase) []Testable {
	testables := []Testable{}
	for _, testCase := range testCases {
		testables = append(testables, testCase)
	}
	return testables
}
