package main

import (
	"fmt"
	"math/cmplx"
)

const Pi = 3.14

var (
	tobe       bool       = false
	maxint     uint64     = 1<<64 - 1
	complexNum complex128 = cmplx.Sqrt(-5 + 2i)
)

func main() {

	fmt.Printf("type: %T value:%v\n", tobe, tobe)
	fmt.Printf("type: %T value:%v\n", maxint, maxint)
	fmt.Printf("type: %T value:%v\n", complexNum, complexNum)
}
