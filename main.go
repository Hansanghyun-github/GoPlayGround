package main

import "fmt"

func TestDefer(str *string) {
	// add world to str
	defer func() {
		*str += " World!"
	}()

	// add hello to str
	*str += "Hello"
}

func main() {
	str := ""
	TestDefer(&str)
	// validate str using assert
	if str != "Hello World!" {
		panic("Test failed")
	}
	fmt.Print("Test passed with: ", str)
}
