package testable

type TestBuilder struct {
	suites []*TestSuite
}

func NewTestBuilder() *TestBuilder {
	builder := new(TestBuilder)
	builder.suites = []*TestSuite{}
	return builder
}

func (this *TestBuilder) AddTestSuite() {
	this.suites = append([]*TestSuite{NewTestSuite(nil, nil)}, this.suites...)
}

func (this *TestBuilder) AddSetUpFunction(setUpFunction func()) {
	suite := this.suites[0].SetSetUpFunction(NewSetUpFunction(setUpFunction))
	this.suites[0] = suite
}

func (this *TestBuilder) AddTestCase(testFunction func()) {
	suite := this.suites[0].AddTest(NewTestCase(testFunction))
	this.suites[0] = suite
}

func (this *TestBuilder) Build() *TestSuite {
	last := this.suites[0]
	for i := 1; i < len(this.suites); i++ {
		suite := this.suites[i].AddTest(last)
		last = suite
	}
	return last
}
