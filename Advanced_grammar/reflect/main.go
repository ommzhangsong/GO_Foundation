package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
	age  int
}

func main() {
	num := 100
	t1 := reflect.TypeOf(num)
	if t1 == reflect.TypeOf(0) {
		fmt.Println(t1)
	}

	st := student{
		name: "zha",
		age:  18,
	}
	t2 := reflect.ValueOf(st)

	if t2.Kind() == reflect.Struct {
		fmt.Println(t2.Kind())
	}

	fmt.Println(t1.Name())
}
