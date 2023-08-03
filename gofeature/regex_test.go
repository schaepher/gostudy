package gostudy

import (
	"fmt"
	"regexp"
	"testing"
)

func TestX(t *testing.T) {
	var httpregex = regexp.MustCompile(`Http\d+`)
	httpregex.FindString("HttpTraffic")
	if true {
		fmt.Println("matched")
	}
}

func TestRegexpMatch(t *testing.T) {
	reg, _ := regexp.Compile("^(aaa)")
	match := reg.MatchString("dest")
	t.Log(match == true)
}

func TestRegexp(t *testing.T) {
	httpReg, _ := regexp.Compile(`(Http|PassHttp|Errcode)((-)?.+)`)
	rs := httpReg.FindStringSubmatch("PassHttp200")
	t.Log(rs[1], rs[2])

	rs2 := httpReg.FindStringSubmatch("Http200")
	t.Log(rs2)

	rs3 := httpReg.FindStringSubmatch("Errcode-200")
	t.Log(rs3)
}