package testable

type Testable interface {
	NextTest() Testable
	Execute()
	HasNextTest() bool
	Descriptions() []string
}
