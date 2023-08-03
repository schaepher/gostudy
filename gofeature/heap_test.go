package gostudy

import (
	"container/heap"
	"fmt"
	"testing"
)

type BandwidthHeap struct {
	BizType string
	data    []float64
	max     int
}

func NewBandwidthHeap() *BandwidthHeap {
	h := new(BandwidthHeap)
	h.data = make([]float64, 0)
	return h
}

func (h BandwidthHeap) Len() int           { return len(h.data) }
func (h BandwidthHeap) Less(i, j int) bool { return h.data[i] < h.data[j] }
func (h *BandwidthHeap) Swap(i, j int)     { h.data[i], h.data[j] = h.data[j], h.data[i] }

func (h *BandwidthHeap) Push(x interface{}) {
	h.data = append(h.data, x.(float64))
}

func (h *BandwidthHeap) Pop() (x interface{}) {
	n := len(h.data)
	x = h.data[n-1]
	h.data = h.data[0 : n-1]
	return x
}

func TestHeap(t *testing.T) {
	h := NewBandwidthHeap()
	h.BizType = "t"
	heap.Push(h, 1.5)
	heap.Push(h, 1.3)
	heap.Push(h, 1.2)

	fmt.Println(heap.Pop(h))
}

func TestHeapInit(t *testing.T) {
	h := NewBandwidthHeap()
	h.data = []float64{3, 5, 1, 9, 7}
	heap.Init(h)
	fmt.Println(h.data)
	fmt.Println(heap.Pop(h))
	fmt.Println(h.data)
}
