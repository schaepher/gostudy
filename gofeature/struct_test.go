package gostudy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStructEqual(t *testing.T) {
	// 比较指针的地址，不会继续解开指针比较
	type structComparePointerAddress struct {
		Name string
		Age  int
		a    *int
	}
	num := 2
	num2 := 2
	basicPointer1 := structComparePointerAddress{
		Name: "a",
		Age:  1,
		a:    &num,
	}
	basicPointer2 := structComparePointerAddress{
		Name: "a",
		Age:  1,
		a:    &num,
	}
	basicPointer3 := structComparePointerAddress{
		Name: "a",
		Age:  1,
		a:    &num2,
	}
	t.Log(basicPointer1 == basicPointer2)
	t.Log(basicPointer1 != basicPointer3)

	// 结构体本身是值类型，会比较它的字段
	type levelTwo struct {
		Name string
	}
	type levelOne struct {
		b levelTwo
	}
	nestedStruct1 := levelOne{
		b: levelTwo{
			Name: "a",
		},
	}
	nestedStruct2 := levelOne{
		b: levelTwo{
			Name: "a",
		},
	}
	t.Log(nestedStruct1 == nestedStruct2)

	// 如果字段是结构体指针，仍然不会继续解开指针去比较结构体
	type levelOnePointer struct {
		b *levelTwo
	}
	nestedStructPointer1 := levelOnePointer{
		b: &levelTwo{
			Name: "a",
		},
	}
	nestedStructPointer2 := levelOnePointer{
		b: &levelTwo{
			Name: "a",
		},
	}
	t.Log(nestedStructPointer1 != nestedStructPointer2)

	// DeepEqual 会解开指针继续比较
	t.Log(reflect.DeepEqual(nestedStructPointer1, nestedStructPointer2))
}

func TestInterfaceUpdate(t *testing.T) {
	type testStr struct {
		a int
	}
	s := &testStr{a: 10}

	modify := func(a *testStr) {
		a = &testStr{a: 20}
	}
	modify(s)

	fmt.Println(s)
}


func TestStructSliceAsParameterForLoopModify(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make([]Data, 0)
	a = append(a, Data{Name: "before"})

	for _, st := range a {
		st.Name = "after"
	}

	if a[0].Name != "before" {
		t.Log(a)
		t.Fatal()
	}
}

func TestStructPointerSliceAsParameterForLoopModify(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make([]*Data, 0)
	a = append(a, &Data{Name: "before"})

	for _, st := range a {
		st.Name = "after"
	}

	if a[0].Name != "after" {
		t.Log(a)
		t.Fatal()
	}
}