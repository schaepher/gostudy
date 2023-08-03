package gostudy

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/panjf2000/ants/v2"
)

func TestAnts(t *testing.T) {
	wg := new(sync.WaitGroup)
	running := int32(1)
	p, _ := ants.NewPoolWithFunc(2, func(i interface{}) {
		defer wg.Done()

		if atomic.LoadInt32(&running) == 0 {
			return
		}

		t.Log(i)
		time.Sleep(2 * time.Second)
	})

	go func() {
		t := time.Tick(time.Second * 5)
		for {
			select {
			case <-t:
				atomic.StoreInt32(&running, 1)
				return
			}
		}
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		err := p.Invoke(i)
		fmt.Println(i, err)
	}

	wg.Wait()
}
