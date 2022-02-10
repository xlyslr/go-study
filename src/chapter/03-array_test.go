package chapter

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	Array()
}

func TestDimenArray(t *testing.T) {
	DimenArray()
}

func TestArrayParam(t *testing.T) {
	var a = [3]int{1, 2, 3}
	fmt.Printf("%#v %p\n", a, &a)
	ArrayParam(a)
}
