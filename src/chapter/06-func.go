package chapter

import "fmt"

func IntSum(x ...int) int {
	var ret int

	for i, _ := range x {
		ret += i
	}
	//for i := 0; i < len(x); i++ {
	//	ret += x[i]
	//}
	return ret
}

type Calc func(int, int) int

func Calc2(x, y int, c Calc) int {
	return c(x, y)
}

func Add(x, y int) int {
	return x + y
}

// 闭包示例
func Adder(x int) func(int) int {
	return func(i int) int {
		x += i
		return x
	}
}

func Defer(x int) int {
	defer fmt.Printf("%d\n", Add(1, 2))
	fmt.Printf("defer\n")
	return x
}

func PanicA() {
	panic("panic in PanicA")
}

func PanicB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("recover in PanicB %e\n", err)
		}
	}()
	PanicA()
}
