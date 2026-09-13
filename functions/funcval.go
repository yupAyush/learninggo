package main

import (
	"fmt"
	"math"
)

// functions can be passed as values to other functions

func compute(fn func(float64, float64) float64) float64 {
	return fn(3.0, 4.0)
}

func soemt() {
	fmt.Println(compute(math.Pow))

}
