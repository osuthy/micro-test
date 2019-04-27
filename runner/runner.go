package runner

type testRunner struct {
	Result string
}

var( TestRunner = testRunner{} )

var TestMethod func() = nil

func Run() {
	TestMethod()
}
