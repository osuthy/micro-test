package testable

type SetUpFunction struct {
	function func()
}

func NewSetUpFunction(function func()) *SetUpFunction {
	return &SetUpFunction { function: function }
}

func (this *SetUpFunction) Execute() {
	this.function()
}
