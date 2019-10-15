package test

import (
	"github.com/osuthy/micro-test/testable"
	"github.com/osuthy/micro-test/runner"
	"github.com/osuthy/micro-test"
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
	runner.SetPrinter(spy)
	return spy
}

func resetSuites() {
	microtest.Suites = []testable.Testable{}
}

