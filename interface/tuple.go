package main

// go stores interace values as a pair of (type, value). The type is the concrete type that implements the interface, and the value is the actual value of that type. This allows Go to support dynamic dispatch, where the method to be called is determined at runtime based on the actual type of the value stored in the interface.

import "fmt"

type noisemaker interface {
	MakeNoise() string
}

type dog struct {
}

type cat struct {
}

func (d dog) MakeNoise() string {
	return "Woof!"
}

func (c cat) MakeNoise() string {
	return "Meow!"
}

func main() {
	var n noisemaker

	n = dog{}
	fmt.Println(n.MakeNoise())

	n = cat{}
	fmt.Print(n) // internally it is tuple
	fmt.Println(n.MakeNoise())

}
