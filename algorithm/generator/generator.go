package main

import (
	"fmt"
	"sync"
)

type Customer struct {
	ID int32
}

type CustomerGenerator struct {
	cus     []Customer
	results map[int32]Customer
	lock    *sync.RWMutex
}

func NewCustomerGenerator() *CustomerGenerator {
	g := new(CustomerGenerator)
	g.cus = make([]Customer, 0)
	g.results = make(map[int32]Customer)
	g.lock = new(sync.RWMutex)
	return g
}

type getCustomerFn func() Customer

func (g *CustomerGenerator) For(c Customer) (fn getCustomerFn) {
	g.cus = append(g.cus, c)

	fn = func() Customer {
		g.lock.Lock()
		defer g.lock.Unlock()
		if len(g.results) == 0 {
			for _, tempC := range g.cus {
				g.results[tempC.ID] = tempC
			}
		}

		return g.results[c.ID]
	}

	return fn
}

func main() {
	c1 := Customer{1}
	c2 := Customer{2}

	g := NewCustomerGenerator()
	getCustomerFns := make([]getCustomerFn, 0)
	getCustomerFns = append(getCustomerFns, g.For(c1))
	getCustomerFns = append(getCustomerFns, g.For(c2))

	for _, getCustomer := range getCustomerFns {
		c := getCustomer()
		fmt.Println(c.ID)
	}
}
