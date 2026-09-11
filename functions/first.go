package main

import "fmt"

func add(x int, y int) int { // this can be writteen as func(x,y int)
	return x + y
}

func first() {
	fmt.Println(add(3, 3))
}
