package gostudy

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getOrderedTenElementSlice() []int {
	s1 := make([]int, 10)
	for i := 0; i < 10; i++ {
		s1[i] = i
	}
	return s1
}

func printIntSlice(name string, s []int) {
	fmt.Print(name + ": ")
	for _, element := range s {
		fmt.Print(element, ",")
	}
	fmt.Println()
}

func TestMiddleSubSliceAppend(t *testing.T) {
	s1 := getOrderedTenElementSlice()

	s2 := s1[5:8]
	s2 = append(s2, 20)

	if s1[8] != 20 {
		log.Fatalf("s1: %v\n", s1)
	}
}

func TestTailSubSliceAppend(t *testing.T) {
	s1 := getOrderedTenElementSlice()
	s3 := s1
	s3 = append(s3, 11)
	printIntSlice("s1", s1)
	printIntSlice("s3", s3)
}

func TestStructSliceModifiedInFunction(t *testing.T) {
	type Data struct {
		Name string
	}

	fn := func(dts []Data) {
		for idx := range dts {
			dts[idx].Name = "test"
		}
	}

	dts := make([]Data, 0)
	dts = append(dts, Data{})
	dts = append(dts, Data{})
	fn(dts)

	for _, dt := range dts {
		fmt.Println(dt)
	}
}

func TestSlicePointerAsParameterToAppend(t *testing.T) {
	a := make([]int, 0)

	doSth := func(b *[]int) {
		*b = append(*b, 1)
	}

	doSth(&a)

	if len(a) != 1 || a[0] != 2 {
		t.Log(a)
		t.Fatal()
	}
}

func TestSliceAsParameterToAppend(t *testing.T) {
	a := make([]int, 0)

	doSth := func(b []int) {
		b = append(b, 1)
	}

	doSth(a)

	if len(a) != 0 {
		t.Log(a)
		t.Fatal()
	}
}

func TestSliceAsParameterToReplace(t *testing.T) {
	a := make([]int, 0)
	a = append(a, 1)

	doSth := func(b []int) {
		b[0] = 2
	}

	doSth(a)

	if a[0] != 2 {
		t.Log(a)
		t.Fatal()
	}
}

func TestStructSliceModifiedInFunctionByIndexShouldSuccess(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make([]Data, 0)
	a = append(a, Data{Name: "before"})

	doSth := func(b []Data) {
		b[0].Name = "after"
	}

	doSth(a)

	if a[0].Name != "after" {
		t.Log(a)
		t.Fatal()
	}
}

func TestStructSliceItemModifiedInFunctionByTempVariableShouldNotAffectOriginSliceItem(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make([]Data, 0)
	a = append(a, Data{Name: "before"})

	doSth := func(b []Data) {
		data := b[0]
		data.Name = "after"
	}

	doSth(a)

	assert.Equal(t, a[0].Name, "before")
}

func TestEmptySlice(t *testing.T) {
	slice := make([]int, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)

	slice = slice[:0]
	assert.Equal(t, 0, len(slice))
}

func TestSliceCut(t *testing.T) {
	s := []string{"A", "B", "C"}

	f := s[:1]
	assert.Equal(t, 3, cap(f))
	assert.Equal(t, &s[0], &f[0])

	u := append(s[:1], s[2:]...)
	assert.Equal(t, &s[0], &u[0])
}
