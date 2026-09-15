package main

import "fmt"

// in go interface are are implemented implicitly, so if a type has all the methods of an interface, it implements that interface automatically.

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func main1() {

	var i I = T{"hello dfn world"}
	i.M()
}
