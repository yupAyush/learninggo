package main

import "fmt"

func index[T comparable](s []T, x T) int { // comparable is used as contraint for == != type means only accept types which are comparable
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {

	si := []int{1, 2, 3, 4, 5}
	fmt.Println(index(si, 3))
	fmt.Println(index(si, 6))

	ss := []string{"a", "b", "c"}
	fmt.Println(index(ss, "b"))
	fmt.Println(index(ss, "d"))

}
