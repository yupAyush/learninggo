package main

import "fmt"

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(x)
}

func something() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	//init or post statments are optional

	// for sum < 100 { // this can be used as while loop
	// 	fmt.Print(sum)
	// }

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Print(sum)
}
