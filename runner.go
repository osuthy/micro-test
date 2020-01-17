package microtest

import (
	. "github.com/osuthy/micro-test/testable"
)

type ConnectionDefinable interface {
	SetConnectionForLocal(tc TC) (TC, error)
	SetConnectionForK8S(tc TC, namespace string) (TC, error)
}

var cdefs []ConnectionDefinable = []ConnectionDefinable{}

func AppendConnectionDefinable(cdef ConnectionDefinable) {
	cdefs = append(cdefs, cdef)
}

func Run() {
	for _, suite := range Suites {
		s := suite
		for {
			executeSuite(s, createTC())
			if !s.HasNextTest() {
				break
			}
			s = s.NextTest()
		}
	}
}

func createTC() TC {
	tc := TC{"differences": &Differences{}}
	for _, cdef := range cdefs {
		tc, _ = cdef.SetConnectionForLocal(tc)
	}
	return tc
}

func executeSuite(suite Testable, tc TC) {
	suite.Execute(TestContext(tc))
	diffs := tc["differences"].(*Differences)
	if !diffs.isEmpty() {
		printer.Println(toFormatedDescription(suite.Descriptions()))
		for _, v := range diffs.slice {
			printer.Println(v)
		}
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
