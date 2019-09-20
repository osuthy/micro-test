package test

import (
	"github.com/ShoichiroKitano/micro_test/testable"
	"github.com/ShoichiroKitano/micro_test/runner"
	"github.com/ShoichiroKitano/micro_test"
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
	micro_test.Suites = []testable.Testable{}
}

