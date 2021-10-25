package main

import (
	"fmt"
)

func returnValues() int {
	var result int
	defer func() {
		result++
	}()
	return result
}

func namedReturnValues() (result int) {
	defer func() {
		result++
	}()
}

func main() {
	fmt.Println(returnValues())
	fmt.Println(namedReturnValues())
}
