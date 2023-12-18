package main

import "fmt"

// Дана структура Human
// (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name string
	age  int
}

func (h *Human) SayName() {
	fmt.Println("My name is ", h.name)
}

type Action struct {
	Human
}

func main() {
	allan := Human{"allan", 22}

	action := Action{allan}

	action.SayName()
}
