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

//errors in golang is are implemented using the error interface. The error interface has a single method, Error(), which returns a string. In this example, we define a custom error type MyError that implements the error interface by providing an Error() method. The run() function returns an instance of MyError, and in the main function, we check for errors and print them if they occur.
