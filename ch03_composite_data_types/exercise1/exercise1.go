package main

import "fmt"

func convertArrayIntoMap(arr [8]int) map[int]int {
	res := make(map[int]int, len(arr))
	for k, v := range arr {
		res[k] = v
	}
	return res
}

func main() {
	fmt.Println(convertArrayIntoMap([8]int{0, 1, 2, 3, 4, 5, 6, 7}))
}