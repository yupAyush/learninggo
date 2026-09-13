package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func soe() {
	for i, v := range pow { // return index and value
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	//skipping index
	for _, value := range pow {
		fmt.Printf("Value: %d\n", value)
	}

}
