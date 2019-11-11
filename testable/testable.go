package testable

type Testable interface {
	NextTest() Testable
	Execute(c TestContext)
	HasNextTest() bool
	Descriptions() []string
}
