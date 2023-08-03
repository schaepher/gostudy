package gostudy

import (
	"fmt"
	"sync"
	"testing"
)

func TestMapEmpty(t *testing.T) {
	m := make(map[int]struct{ Name string })
	t.Log(m[1].Name)
}

func TestMapAsParameterToModify(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make(map[int]Data, 0)
	a[0] = Data{Name: "before"}

	doSth := func(b map[int]Data) {
		data := b[0]
		data.Name = "after"
	}

	doSth(a)

	if a[0].Name != "before" {
		t.Log(a)
		t.Fatal()
	}
}

func TestMapAsParameterToModifyDirect(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make(map[int]Data, 0)
	a[0] = Data{Name: "before"}

	doSth := func(b map[int]Data) {
		// b[0].Name = "after" // error
	}

	doSth(a)

	if a[0].Name != "before" {
		t.Log(a)
		t.Fatal()
	}
}

func TestMapAsParameterToModifyDirectPointer(t *testing.T) {
	type Data struct {
		Name string
	}
	a := make(map[int]*Data, 0)
	a[0] = &Data{Name: "before"}

	doSth := func(b map[int]*Data) {
		b[0].Name = "after"
	}

	doSth(a)

	if a[0].Name != "after" {
		t.Log(a)
		t.Fatal()
	}
}

func TestMapDelete(t *testing.T) {
	type Value struct {
		name map[int64]int64
	}

	v := &Value{name: make(map[int64]int64)}

	lock := new(sync.Mutex)
	lock.Lock()

	wg := new(sync.WaitGroup)
	wg.Add(100000)

	go func() {
		for i := int64(0); i < 100000; i++ {
			v.name = map[int64]int64{
				i: i,
			}
			wg.Done()
		}
	}()

	go func() {
		for {
			for key, value := range v.name {
				fmt.Println(key, value)
				lock.Unlock()
			}
		}
	}()

	wg.Wait()
}
