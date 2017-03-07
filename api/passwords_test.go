package main

import "fmt"

func ExampleHash() {
	// s := CleanJSON("}\"")
	// fmt.Println(s)

	// s = CleanJSON("\"{")
	// fmt.Println(s)

	// s = CleanJSON("\\\"")
	// fmt.Println(s)

	// // Output: }
	// // {
	// // "

	password := "secret"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	//fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)

	// Output: Password: secret
	// Match:    true

}
