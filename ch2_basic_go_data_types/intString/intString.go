package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 100
	input := strconv.Itoa(n)
	fmt.Printf("strconv.Itoa() %s of type %T\n", input, input)
	input = strconv.FormatInt(int64(n), 10)
	fmt.Printf("strconv.FormatInt() %s of type %T\n", input, input)
	input = string(n)
	fmt.Printf("string() %s of type %T\n", input, input)

	r := 'A'
	fmt.Println(r)
	fmt.Printf("%T\n", r)
	fmt.Println(string(r))
}