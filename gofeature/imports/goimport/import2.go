package goimport

import (
	"fmt"
	"gofeature/imports/goimport2"
)


func init() {
	fmt.Println("import 2")
}

func Call2() {
	fmt.Println("call 2")
	goimport2.Call4()
}