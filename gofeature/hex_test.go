package gostudy

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	x, err := hex.DecodeString("ae3d6031503c5aadabe0213902fb3214d1a5234a")
	if err != nil {
		return
	}
	ioutil.WriteFile("a.sig", x, os.ModePerm)
}
