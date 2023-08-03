package gostudy

import (
	"fmt"
	"testing"
	"time"
)

func TestMapAssign(t *testing.T) {
	type Value struct {
		Num int32
	}
	m := map[int32]Value{
		1: {1},
	}
	fmt.Println(m[1].Num)

	// UnaddressableFieldAssign occurs when trying to assign to a struct field
	// m[1].Num = 2
}

func TestSliceAssign(t *testing.T) {
	type Value struct {
		Num int32
	}
	s := []Value{
		{1},
	}

	s[0].Num = 2

	if s[0].Num != 2 {
		t.Error("not equal")
	}
}

func talk(msg string, sleep int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}()
	return ch
}

func fanIn(input1, input2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			select {
			case ch <- <-input1:
			case ch <- <-input2:
			}
		}
	}()
	return ch
}

func TestChanDeadLock(t *testing.T) {
	ch := fanIn(talk("A", 10), talk("B", 1000))
	for i := 0; i < 10; i++ {
		fmt.Printf("%q\n", <-ch)
	}
}
