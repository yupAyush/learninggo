package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex //[key]value

func main() {
	m = make(map[string]Vertex) // make initializes the map so it can be assigned to
	m["google"] = Vertex{
		2344.2244, -24442.2344,
	}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m["hello world"])
	fmt.Println(m["google"])

	v, ok := m["google"]
	if ok {
		fmt.Println("Value is present", v)
	} else {
		fmt.Println("Value is not present")
	}

}
