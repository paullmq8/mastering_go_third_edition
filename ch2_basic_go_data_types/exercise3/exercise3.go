package main

import "fmt"

func concatenateSlices(slc1, slc2 []int) [8]int {
	slc1 = append(slc1, slc2...)
	if len(slc1) >= 8 {
		return *(*[8]int)(slc1)
	}
	return [8]int{}
}

func main() {
	res := concatenateSlices([]int{0, 1, 2, 3}, []int{4, 5, 6, 7, 8})
	fmt.Println(res)
	res = concatenateSlices([]int{0, 1}, []int{4, 5, 6, 7})
	fmt.Println(res)

	var aMap map[string]int
	aMap = nil
	fmt.Println(aMap)
}
