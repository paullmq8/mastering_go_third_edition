package main

import "fmt"

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
}

func appendSlice(arr []int) []int {
	fmt.Printf("%p\n", &arr)
	arr = append(arr, 4)
	arr = append(arr, 5)
	return arr
}