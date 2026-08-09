package main

import "fmt"

func pointers() {

	age := 29 // regular var
	var agePointer *int
	agePointer = &age

	fmt.Println("Age: ", *agePointer)
}
