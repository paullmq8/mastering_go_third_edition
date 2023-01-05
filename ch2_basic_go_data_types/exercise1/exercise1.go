package main

import "fmt"

func concatenateArrays(arr1, arr2 [4]int) []int {
	result := make([]int, 0, len(arr1) + len(arr2))
	result = append(result, arr1[:]...)
	result = append(result, arr2[:]...)
	return result
}

func Foo() [32]byte {
    return [32]byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
}

func Bar(b []byte) {
    fmt.Println(string(b))
}

func main() {
	totalSlice1 := concatenateArrays([4]int{0, 1, 2, 3}, [4]int{4, 5, 6, 7})
	fmt.Println(totalSlice1)
	totalSlice2 := concatenateArrays([4]int{}, [4]int{})
	fmt.Println(totalSlice2)

	x := Foo()
    Bar(x[:])
}
