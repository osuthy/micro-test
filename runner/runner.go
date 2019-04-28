package runner

type testRunner struct {
	Result string
}

var TestRunner testRunner = testRunner{}
var	TestFunctions []func() = []func(){}


func Run() {
	for _, TestFunction := range TestFunctions {
		TestFunction()
	}
}
