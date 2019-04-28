package testable

type SetUpFunction struct {
	function func()
}

func NewSetUpFunction(function func()) *SetUpFunction {
	setUpFunction := new(SetUpFunction)
	setUpFunction.function = function
	return setUpFunction
}

func (this *SetUpFunction) Execute() {
	this.function()
}
