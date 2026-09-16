package main

import "fmt"

//type assertion is used to retrieve the dynamic value of an interface. It allows you to access the underlying concrete value stored in an interface variable. The syntax for type assertion is as follows:

func main7() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic: interface conversion: interface {} is string, not float64
	fmt.Println(f)

}
