package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(`This is 
	raw 
	string`)
	re := regexp.MustCompile(`\w`)
	s := "s_s"
	if re.Match([]byte(s)) {
		fmt.Println(s + " is word")
	} else {
		fmt.Println(s + " is not word")
	}
}