package main

import "fmt"

// how methods works in go , even if go doesnt have classes, it has methods. Methods are functions that are associated with a particular type. In Go, you can define methods on any type you create, including structs and custom types.
// they are defined using a receiver function , which is a special parameter that allows the method to access the data of the type it is associated with. The receiver function is specified before the method name and is enclosed in parentheses. The receiver can be a value receiver or a pointer receiver, depending on whether you want to modify the original value or work with a copy of it.

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 { // when we pass simple type in the receiver function , it passes by value and we cannot modify the original value of the struct. If we want to modify the original value of the struct, we can use a pointer receiver instead.
	return r.width * r.height
}

// (r *rectangle) area() float64 { // when we pass pointer type in the receiver function , it passes by reference and we can modify the original value of the struct.
// 	return r.width * r.height
// }

// methods can work on non struct types as well
type number int

func (n number) square() number {
	return n * n
}

// methods can take value and pointer receivers, and they can be called on both values and pointers of the type. When you call a method on a value, Go automatically creates a copy of the value and passes it to the method. When you call a method on a pointer, Go automatically dereferences the pointer and passes the value to the method.
// while function can take only those arguments which are defined iin there function signature
// we should use pointers receiver when struct is large and we want to avoid copying the entire struct, or when we want to modify the original value of the struct. We should use value receiver when the struct is small and we don't need to modify the original value of the struct.

func main() {
	r := rectangle{width: 5, height: 10}
	fmt.Println(r.area())

	n := number(4)
	fmt.Println(n.square())

}
