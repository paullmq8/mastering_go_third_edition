package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 1
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println(t, v)
}