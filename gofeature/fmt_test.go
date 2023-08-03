package gostudy

import (
	"fmt"
	"testing"
)

func TestFormat(t *testing.T) {
	fmt.Printf("%0.2f", float64(11.8999))
}

func TestPadding(t *testing.T) {
	s := fmt.Sprintf("%04d", 12)
	t.Log(s)
}


func TestPrintByte(t *testing.T) {
	fmt.Printf("%c", 'A')
}
