package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a rune = 'c'
	var b int32 = 'd'
	if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a > b")
	}

	aString := "Hello World! €"
	fmt.Printf("%T\n", aString[0]) // uint8
	fmt.Println("First character", string(aString[0])) // H

	// Runes
	// A rune
	r := '€'
	fmt.Printf("%T\n", r)
	fmt.Println("As an int32 value:", r)
	// Convert Runes to text
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)
	// Print an existing string as runes
	for _, v := range aString {
		fmt.Printf("%T %c %x\n", v, v, v) // %T is always int32 here
	}
	fmt.Println()
	r2 := string(0x20ac)
	fmt.Printf("%s\n", r2)
	r3 := strconv.FormatInt(0x20ac, 16)
	r4 := strconv.Itoa(0x20ac)
	r5, _ := strconv.Atoi("123")
	fmt.Printf("%s %s %s %d\n", r2, r3, r4, r5)
}