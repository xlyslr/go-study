package chapter

import "fmt"

func Array() {
	var a [3]int
	fmt.Printf("%#v\n", a)

	var b = [3]int{}
	fmt.Printf("%#v\n", b)

	var c = [3]int{1, 2}
	fmt.Printf("%#v\n", c)

	var d = [...]int{1, 2, 3}
	fmt.Printf("%#v\n", d)

	var e = [...]int{1: 1, 5: 2, 7: 3}
	fmt.Printf("%#v\n", e)

	fmt.Printf("%t\n", a == b)
}

func DimenArray() {
	var a [3][3]int
	fmt.Printf("%#v\n", a)

	var b = [3][3]int{{1}}
	fmt.Printf("%#v\n", b)

	fmt.Printf("%#v\n", b[2])
	fmt.Printf("%#v\n", b[2][1])
}

func ArrayParam(a [3]int) {
	fmt.Printf("%#v %p\n", a, &a)
}
