package main

import "fmt"

type person struct {
	Name string
	age  int
}

func (p person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.age)

}

func main() {
	a := person{"Arthur Dent", 42}
	z := person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
