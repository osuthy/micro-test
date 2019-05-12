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
	this.suites = append([]*TestSuite{NewTestSuite()}, this.suites...)
}

func (this *TestBuilder) AddSetUpFunction(setUpFunction func()) {
	this.suites[0].setUpFunction = NewSetUpFunction(setUpFunction)
}

func (this *TestBuilder) AddTestCase(testFunction func()) {
	this.suites[0].Add(NewTestCase(testFunction))
}

func (this *TestBuilder) Build() Testable {
	last := this.suites[0]
	for i := 1; i < len(this.suites); i++ {
		this.suites[i].Add(last)
		last = this.suites[i]
	}
	return last
}
