package main

import "fmt"

func main() {
	var a [5]int
	a[0] = 45
	a[1] = 90
	fmt.Println(a)
	fmt.Println(a[0], a[1])
	primes := [8]int{2, 3, 5, 7, 11, 13, 17, 19}
	fmt.Println(primes)
	b := primes[1:5]
	b[0] = 12121212

	fmt.Println(primes[1:5])
	fmt.Println(primes)
	fmt.Println(len(b))
	fmt.Println(cap(b))

}
