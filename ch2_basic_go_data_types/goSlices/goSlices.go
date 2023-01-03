package main

import "fmt"

func main() {
	arr := [4]int{}
	fmt.Println(arr[1], arr[3])
	aSlice := make([]float64, 0)
	fmt.Printf("%T %v\n", aSlice, aSlice)

	twoDSlice := make([][]int, 2)
	fmt.Printf("%T %d %d %v %v\n", twoDSlice, len(twoDSlice), len(twoDSlice[0]), twoDSlice[0], twoDSlice[1])

	fmt.Println("=================")
	twoDSlice[0] = []int{1, 2, 3}
	twoDSlice[1] = []int{4, 5, 6, 7, 8}
	fmt.Printf("%T %d %d %v %v\n", twoDSlice, len(twoDSlice), len(twoDSlice[0]), twoDSlice[0], twoDSlice[1])
	fmt.Println("=================")

	// Create an empty slice
	aSlice2 := []float64{}
	// Both length and capacity are 0 because slice is empty
	fmt.Println(aSlice2, len(aSlice2), cap(aSlice2))
	// Add elements to a slice
	aSlice2 = append(aSlice2, 1234.56)
	aSlice2 = append(aSlice2, -34.0)
	fmt.Println(aSlice2, "with length", len(aSlice2))

	// A slice with length 4
	t := make([]int, 4)
	fmt.Println(t[3])
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	// Now you will need to use append
	t = append(t, -5)
	fmt.Println(t)

	// A 2D slice
	// You can have as many dimensions as needed
	twoD := [][]int{{1,2,3}, {4,5,6}}
	// Visiting all elements of a 2D slice
	// with a double for loop
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1,2,3,4}
	make2D[1] = []int{-1,-2,-3,-4}
	fmt.Println(make2D)
}