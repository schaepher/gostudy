package gostudy

import (
	"fmt"
	"testing"
)

type Item struct {
	Prev  *Item
	Next  *Item
	Key   Key
	Value Value
}

type Key struct {
	K int32
}

type Value struct {
	V int32
}

type LRUList struct {
	cache   map[Key]*Item
	head    *Item
	tail    *Item
	count   int
	maxSize int
}

func NewLRUList(maxSize int) *LRUList {
	h := &Item{}
	return &LRUList{
		cache:   make(map[Key]*Item),
		head:    h,
		tail:    h,
		maxSize: maxSize,
	}
}

func (l *LRUList) hot(item *Item) {
	first := l.head.Next
	if first == item {
		return
	}

	item.Prev.Next = item.Next
	item.Next.Prev = item.Prev

	l.head.Next = item
	item.Prev = l.head
	item.Next = first
}

func (l *LRUList) Put(k Key, v Value) {
	if item, ok := l.cache[k]; ok {
		l.hot(item)
		item.Value = v
		return
	}

	item := &Item{Key: k, Value: v}
	l.cache[k] = item
	if l.head.Next != nil {
		l.head.Next.Prev = item
	}
	item.Next = l.head.Next
	item.Prev = l.head
	l.head.Next = item
	l.count++

	if item.Next == nil {
		l.tail = item
	}

	if l.count > l.maxSize {
		prev := l.tail.Prev
		prev.Next = nil
		delete(l.cache, l.tail.Key)
		l.tail = prev
		l.count--
	}
}

func (l *LRUList) Get(k Key) (Value, bool) {
	item, ok := l.cache[k]
	if !ok {
		return Value{}, false
	}

	l.hot(item)

	return item.Value, true
}

func (l *LRUList) PrintAll() {
	cur := l.head
	for cur.Next != nil {
		cur = cur.Next
		fmt.Println(cur.Value.V)
	}
}

func TestLRU(t *testing.T) {
	l := NewLRUList(5)
	l.Put(Key{1}, Value{1})
	l.Put(Key{2}, Value{2})
	l.Put(Key{3}, Value{3})
	l.Put(Key{4}, Value{4})
	l.Put(Key{5}, Value{5})
	l.Put(Key{6}, Value{6})
	l.Put(Key{6}, Value{6})
	l.Get(Key{3})
	l.Put(Key{1}, Value{1})

	l.PrintAll()
}
