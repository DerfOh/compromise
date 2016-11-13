package main

import "fmt"

func ExampleCleanJSON() {
	s := CleanJSON("}\"")
	fmt.Println(s)

	s = CleanJSON("\"{")
	fmt.Println(s)

	s = CleanJSON("\\\"")
	fmt.Println(s)

	// Output: }
	// {
	// "
}
