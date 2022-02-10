package chapter

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var c Calc
	c = Add
	fmt.Printf("%d\n", c(1, 2))
}

func TestCalc2(t *testing.T) {
	fmt.Printf("%d\n", Calc2(1, 2, Add))
}

func TestAdder(t *testing.T) {
	a := Adder(10)
	fmt.Printf("%d\n", a(10))
	fmt.Printf("%d\n", a(20))
	fmt.Printf("%d\n", a(30))
}

func TestDefer(t *testing.T) {
	fmt.Printf("%d\n", Defer(1))
}

func TestPanicB(t *testing.T) {
	PanicB()
}
