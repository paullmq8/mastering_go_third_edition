package main

import (
	"fmt"
	"os"
)

func main() {
	// Get user input
	fmt.Printf("Please give me your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)

	a := 1
	fmt.Println(a, -a)
	b := -1
	if b < 0 {
		b = -b
	}
	fmt.Println(b)

	arr1 := []int{1, 2, 3}
	arr1 = appendSlice(arr1)
	fmt.Printf("%p\n", &arr1)
	fmt.Println(arr1)

	// go run input.go 'a b c'
	// /var/folders/2w/srvhlv8d13l_l1dxlcs1_2y40000gp/T/go-build3860030454/b001/exe/input a b c
	// go run input.go "a b c"
	// /var/folders/2w/srvhlv8d13l_l1dxlcs1_2y40000gp/T/go-build3920247801/b001/exe/input a b c
	// ./input abc
	// ./input abc
	fmt.Println(os.Args[0], os.Args[1])
}

func appendSlice(arr []int) []int {
	fmt.Printf("%p\n", &arr)
	arr = append(arr, 4)
	arr = append(arr, 5)
	return arr
}