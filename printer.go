package microtest

import (
	"github.com/k0kubun/pp"
)

type Printable interface {
	Println(str string)
}

type Printer struct{}

var printer Printable = Printer{}

func (this Printer) Println(str string) {
	pp.Println(str)
}

// for auto test
func SetPrinter(p Printable) {
	printer = p
}
