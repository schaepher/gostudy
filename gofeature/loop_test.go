package gostudy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForAssignAddressShouldAllSameAsLastOne(t *testing.T) {
	type St struct {
		Name string
	}

	testFor := map[int]St{
		1: {"a"},
		2: {"b"},
		3: {"c"},
	}

	rs := make([]*St, 0, len(testFor))
	var lastName string
	for _, b := range testFor {
		rs = append(rs, &b)
		lastName = b.Name
	}

	assert.Equal(t, lastName, rs[0].Name)
	assert.Equal(t, lastName, rs[1].Name)
	assert.Equal(t, lastName, rs[2].Name)
	assert.Equal(t, rs[2], rs[0])
	assert.Equal(t, rs[2], rs[1])
	assert.Equal(t, rs[0], rs[1])
}

func TestForAssignStructShouldKeepOriginData(t *testing.T) {
	type St struct {
		Name string
	}

	testFor := map[int]St{
		1: {"a"},
		2: {"b"},
		3: {"c"},
	}

	rs := make([]St, 0, len(testFor))
	var lastName string
	for _, b := range testFor {
		rs = append(rs, b)
		lastName = b.Name
	}

	assert.NotEqual(t, lastName, rs[0].Name)
	assert.NotEqual(t, lastName, rs[1].Name)
	assert.Equal(t, lastName, rs[2].Name)
	assert.NotEqual(t, rs[2], rs[0])
	assert.NotEqual(t, rs[2], rs[1])
	assert.NotEqual(t, rs[0], rs[1])
}
