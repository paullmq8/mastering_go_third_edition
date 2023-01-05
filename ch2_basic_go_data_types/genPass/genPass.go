package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var MIN = 0
var MAX = 94

func random(min, max int) int {
	return rand.Intn(max - min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func main() {

	startChar := "!"
	char := '!'
	fmt.Println(startChar[0])
	fmt.Println(char)
	fmt.Println(string(startChar[0]))
	fmt.Println("=========")
	myRand := random(MIN, MAX)
	fmt.Println(myRand, byte(myRand))
	fmt.Println(string(startChar[0] + byte(myRand)))
	fmt.Println("=========")

	var LENGTH int64 = 8

	arguments := os.Args
	switch len(arguments) {
	case 2:
		t, err := strconv.ParseInt(arguments[1], 10, 64)
		if err == nil {
			LENGTH = t
		}
		if LENGTH <= 0 {
			LENGTH = 8
		}
	default:
		fmt.Println("Using default values...")
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	fmt.Println(getString(LENGTH))
	fmt.Println("========")
	fmt.Printf("%T %b\n", startChar[0], startChar[0])
	fmt.Println(strconv.ParseInt("1234", 16, 64))
}