package runner

import (
	"github.com/k0kubun/pp"
)

type Printable interface {
	Println(str string)
}

type Printer struct {}

var printer Printable = Printer{}

// for auto test
func SetPrinter(p Printable) {
	printer = p
}

func (this Printer) Println(str string) {
	pp.Println(str)
}
