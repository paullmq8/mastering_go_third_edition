package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	re := regexp.MustCompile(`^[-+]?\d+$`)
	return re.Match(t)
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: <utility> string string string...")
		return
	}

	s := arguments[1:]
	trueCount, falseCount := 0, 0
	for _, v := range s {
		ret := matchInt(v)
		if ret {
			trueCount++
		} else {
			falseCount++
		}
		fmt.Println(v, ret)
	}
	fmt.Println("true count:", trueCount, ", false count:", falseCount)
}