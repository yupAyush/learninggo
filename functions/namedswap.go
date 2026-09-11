package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum / 2
	y = sum - x
	return
}

func something2() {
	fmt.Println(split(17))
}
