package testable

type Testable interface {
	Execute()
	HasUnexecutedTest() bool
}
