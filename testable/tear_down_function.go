package testable

type TearDownFunction struct {
	function func()
}

func NewTearDownFunction(function func()) *TearDownFunction {
	return &TearDownFunction{function: function}
}

func (this *TearDownFunction) Execute() {
	this.function()
}
