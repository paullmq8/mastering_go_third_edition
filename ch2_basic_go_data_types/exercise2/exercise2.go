package main

import "fmt"

func concatenateTwoArrays(arr1, arr2 [4]int) [8]int {
	res := make([]int, 0, len(arr1) + len(arr2))
	res = append(res, arr1[:]...)
	res = append(res, arr2[:]...)
	ret := (*[8]int)(res)
	// ret1 := (*[5]int)(res)
	// fmt.Println(ret1)
	return *ret
}

func main() {
	result := concatenateTwoArrays([4]int{0, 1, 2, 3}, [4]int{4, 5, 6, 7})
	fmt.Println(result)
}