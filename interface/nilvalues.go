package main

import "fmt"

// in go there is no null pointer, instead there is nil value which can be assigned to pointers, interfaces, maps, slices, channels and function types. A nil interface value holds neither value nor concrete type.

type i interface {
	M()
}

type t struct {
	s string
}

func (t *t) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.s)
}

func main4() {
	var i i
	i = (*t)(nil)
	i.M()

	i = &t{"hello world"}
	i.M()

}

func describe(x i) {
	fmt.Printf("(%v, %T)\n", x, x)
}
