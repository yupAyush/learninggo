package main

import (
	"fmt"
	"runtime"
	"time"
)

func son() {

	fmt.Println("when's saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today is saturday")
	case today + 1:
		fmt.Println("tomorrow is saturday")
	case today + 2:
		fmt.Println("saturday is two days away")
	default:
		fmt.Println("saturday is too far away")
	}

	fmt.Println("go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macos")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s", os)
	}
}
