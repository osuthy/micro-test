package runner

import (
	"github.com/ShoichiroKitano/micro_test"
	"github.com/ShoichiroKitano/micro_test/testable"
)

var Diffs *Differences = &Differences{}

func Run() {
	for _, suite := range micro_test.Suites {
		executeSuite(suite)
	}
}

func executeSuite(suite testable.Testable) {
	s := suite
	for {
		s.Execute()
		if !Diffs.isEmpty() {
			printer.Println(toFormatedDescription(s.Descriptions()))
			for _, v := range Diffs.slice {
				printer.Println(v)
			}
		}
		Diffs = &Differences{}

		if !s.HasNextTest() {
			return
		}

		s = s.NextTest()
	}
}

func toFormatedDescription(descriptions []string) string {
	desc := ""
	for i, v := range descriptions {
		if i == 0 {
			desc = v
		} else if v != "" {
			desc = desc + " " + v
		}
	}
	return desc
}
