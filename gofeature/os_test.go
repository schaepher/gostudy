package gostudy

import (
	"os"
	"testing"
)

func TestGetHostNameTest(t *testing.T) {
	name, err := os.Hostname()
	t.Log(name, err)
}
