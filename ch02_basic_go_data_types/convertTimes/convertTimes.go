package main

import (
	"fmt"
	"os"
	"time"
)

// go run convertTimes.go "14 December 2020 19:20 EET"
// go run convertTimes.go "14 December 2020 20:00 UTC"
// go run convertTimes.go "14 December 2020 25:00 EET"
func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need input in __02 January 2006 15:04 MST__ format!")
		return
	}

	input := os.Args[1]
	now, err := time.Parse("02 January 2006 15:04 MST", input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Local time
	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current Location: %s\n", now.In(loc))

	// NY
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	// London
	loc, _ = time.LoadLocation("Europe/London")
	fmt.Printf("London Time: %s\n", now.In(loc))

	// Tokyo
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo Time: %s\n", now.In(loc))
}