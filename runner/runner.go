package runner

import (
	"github.com/ShoichiroKitano/micro_test"
)

type testRunner struct {
	Result string
}

var Queue *DifferenceQueue = &DifferenceQueue{}

var TestRunner testRunner = testRunner{}

func Run() {
	for _, suite := range micro_test.Suites {
		s := suite
		for {
			s.Execute()
			if !Queue.isEmpty() {
				desc := ""
				for i, v := range s.Descriptions() {
					if i == 0 {
						desc = v
					} else {
						desc = desc + " " + v
					}
				}
				printer.Println(desc)

				for _, v := range Queue.queue {
					printer.Println(v)
				}
			}
			Queue = &DifferenceQueue{}
			if !s.HasNextTest() {
				break
			}
			s = s.NextTest()
		}
	}
}

