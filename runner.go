package microtest

import (
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
		diffs := &Differences{}
		tc["differences"] = diffs
		s.Execute(TestContext(tc))
		if !diffs.isEmpty() {
			printer.Println(toFormatedDescription(s.Descriptions()))
			for _, v := range diffs.slice {
				printer.Println(v)
			}
		}
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
