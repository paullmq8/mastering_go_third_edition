package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	invalid := make([]string, 0)
	var nInts, nFloats, total int
	for _, k := range arguments[1:] {
		// Is it an integer?
		_, err := strconv.Atoi(k)
		if err == nil {
			nInts++
			total++
			continue
		}
		// Is it a float?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			nFloats++
			total++
			continue
		}
		// Then it is invalid
		invalid = append(invalid, k)
	}
	if len(invalid) > nInts + nFloats {
		fmt.Println("invalids are greater than ints and floats.")
	} else {
		fmt.Printf("#read: %d #ints: %d #floats: %d", nInts + nFloats, nInts, nFloats)
	}
}