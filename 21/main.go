package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type contract interface {
	doSomething()
}

type Human struct {
	name string
	age  int
}

func (h *Human) eat() {
	fmt.Println("eating")
}

func (h *Human) sleep() {
	fmt.Println("sleaping")
}

type humanAdapter struct {
	human *Human
}

func (adapter *humanAdapter) doSomething() {
	adapter.human.eat()
	adapter.human.sleep()
}

func run(inter contract) {
	inter.doSomething()
}

func main() {
	human := Human{"allan", 5}
	adapter := humanAdapter{&human}

	run(&adapter)
}
