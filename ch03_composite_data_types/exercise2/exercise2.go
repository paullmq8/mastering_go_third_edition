package main

import "fmt"

func convertMapIntoTwoSlices(input map[string]int) (slice1 []string, slice2 []int) {
	slice1 = make([]string, 0, len(input))
	slice2 = make([]int, 0, len(input))
	for k, v := range input {
		slice1 = append(slice1, k)
		slice2 = append(slice2, v)
	}
	return slice1, slice2
}

func main() {
	res := make(map[string]int, 3)
	res["aaa"] = 1
	res["bbb"] = 2
	res["ccc"] = 3
	slice1, slice2 := convertMapIntoTwoSlices(res)
	fmt.Println(slice1)
	fmt.Println(slice2)
}