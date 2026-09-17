package main

import "fmt"

type MyError struct {
	When string
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %s, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		When: "now",
		What: "it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
