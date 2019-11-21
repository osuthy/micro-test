package runner

import (
	. "github.com/osuthy/micro-test"
	. "github.com/osuthy/micro-test/testable"
)

type ConnectionDefinable interface {
	SetConnectionForLocal(tc TC) (TC)
	SetConnectionForK8S(tc TC, namespace string) (TC)
}

var cdefs []ConnectionDefinable = []ConnectionDefinable{}

func AppendConnectionDefinable(cdef ConnectionDefinable) {
	cdefs = append(cdefs, cdef)
}

var Diffs *Differences = &Differences{}

func Run() {
	for _, suite := range Suites {
		executeSuite(suite)
	}
}

func executeSuite(suite Testable) {
	tc := TC{}
	for _, cdef := range cdefs {
		tc = cdef.SetConnectionForLocal(tc)
	}
	s := suite
	for {
		s.Execute(TestContext(tc))
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
