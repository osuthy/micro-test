package runner

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/testable"
)

var Diffs *Differences = &Differences{}

func Run() {
	for _, suite := range Suites {
		executeSuite(suite)
	}
}

func executeSuite(suite Testable) {
	s := suite
	for {
		s.Execute(TestContext{})
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
