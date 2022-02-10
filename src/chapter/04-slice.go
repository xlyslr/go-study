package chapter

import "fmt"

func Slice() {
	var a []int
	fmt.Printf("%#v\n", a)

	b := make([]int, 2)
	fmt.Printf("%#v\n", b)
}
