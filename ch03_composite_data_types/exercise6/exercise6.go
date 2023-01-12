package main

import (
	"fmt"
	"os"
)

type param struct {
	Idx int
	Val string
}

func main() {
	slc := make([]param, 0, len(os.Args))
	for k, v := range os.Args {
		slc = append(slc, param{Idx: k, Val: v})
	}
	fmt.Println("slc:", slc)
}