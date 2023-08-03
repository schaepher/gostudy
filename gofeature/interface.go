package gostudy

import (
	"testing"
)

func TestInterfaceCompare(t *testing.T) {
	var a interface{}
	var b interface{}

	c := int64(1)
	d := int64(1)

	a = c
	b = d

	if a == b {
		t.Log("equal")
	}
}
