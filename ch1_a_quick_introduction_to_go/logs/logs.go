package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(len(os.Args))
	if len(os.Args) != 1 {
		log.Fatal("Fatal: Hello World!")
	}
	log.Panic("Panic: Hello World!")
}