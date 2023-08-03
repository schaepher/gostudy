package main

import (
	"os"
	"strconv"
	"strings"
)


// usage: 
// 1. go run . "print a"
// 2. go run . "add 1 2"
func main() {
	script := os.Args[1]

	scripts := strings.Split(script, " ")
	if scripts[0] == "print" {
		println(scripts[1])
	} else if scripts[0] == "add" {
		a, _ := strconv.Atoi(scripts[1])
		b, _ := strconv.Atoi(scripts[2])
		println(a+b)
	}
}