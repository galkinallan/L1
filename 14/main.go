package main

import (
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип
// переменной: int, string, bool, channel из переменной типа interface{}.

func typeOf(x any) string {
	return reflect.TypeOf(x).String()
}

func main() {
	num := 5
	str := "hello"
	fl := 1.1

	println(typeOf(num))
	println(typeOf(str))
	println(typeOf(fl))
}
