package main

import "fmt"

func printEntry() {
	type Entry struct {
		Name string
		Surname string
		Year int
	}
	anEntry := Entry{
		Name: "aaa",
		Surname: "bbb",
		Year: 16,
	}
	fmt.Println(anEntry)
}

func main() {
	printEntry()
/* 	anEntry := Entry{
		Name: "aaa",
		Surname: "bbb",
		Year: 16,
	}
	fmt.Println(anEntry) */
}