package test

import (
	"github.com/osuthy/micro-test"
	"github.com/osuthy/micro-test/testable"
)

type PrinterSpy struct {
	results []string
}

func NewPrinterSpy() *PrinterSpy {
	return new(PrinterSpy)
}

func (this *PrinterSpy) Println(str string) {
	this.results = append(this.results, str)
}

func setUpSpy() *PrinterSpy {
	spy := NewPrinterSpy()
	microtest.SetPrinter(spy)
	return spy
}

func resetSuites() {
	microtest.Suites = []testable.Testable{}
}
