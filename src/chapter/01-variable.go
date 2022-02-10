package chapter

import "fmt"

// 声明
func Statement() {

	var v1 int
	v1 = 1
	fmt.Printf("%d\n", v1)
	// 批量声明
	var (
		v2 int
		v3 string
	)
	v2 = 1
	v3 = "2"
	fmt.Printf("%d %s\n", v2, v3)

	// 整型数组
	var v4 [2]int32
	v4 = [...]int32{1, 2}

	fmt.Printf("%v\n", v4)

	// 结构体
	var v5 struct {
		a int
		b string
	}
	v5.a = 1
	v5.b = "2"
	fmt.Printf("%v\n", v5)
	// 整型指针
	var v6 *int
	fmt.Printf("%p\n", v6)

	// map字典 key为string类型 value为int类型
	var v7 map[string]int
	v7 = make(map[string]int, 8)
	v7["a"] = 1
	v7["b"] = 2

	fmt.Printf("%v\n", v7)

	// 函数 入参 出参
	var v8 func(a int) int
	v8 = func(a int) int {
		fmt.Printf("%d\n", a)
		return a
	}
	v8(1)
}

// 初始化
func Initialize() {
	var v1 int = 1
	var v2 = 1
	v3 := 1
	fmt.Printf("%d %d %d\n", v1, v2, v3)
}

//== 多重赋值
func MultipleAssign() {
	i := 1
	j := 2
	// 交换变量值，不需要中间变量
	i, j = j, i
}

//== 匿名变量 _，不想要的变量可以用_忽略
func GetName() (a, b string) {
	return "1", "2"
}

func Test1() {
	name, _ := GetName()
	fmt.Println(name)
}

// **变量作用域：函数体外声明的可以被认为是全局变量，在函数体内声明的是局部变量，函数的参数和返回值变量也是局部变量**
// 常量通过`const`来定义，只能是数值类型、布尔类型、字符串类型等标量类型
const Pi = 3.141592653
const zero = 0.0
const (
	size = 1024
	eof  = -1
)

// true和false实际是内置常量
const (
	true1  = true  //const true bool = 0 == 0
	false1 = false //const false bool = 0 != 0
)

// 内置常量iota，在每次const出现时重置为0，然后每使用一次自增1
const ( // iota 被重置为 0
	c0 = iota // c0 = 0
	c1 = iota // c1 = 1
	c2 = iota // c2 = 2
)

const (
	u = iota * 2 // u = 0
	v = iota * 2 // v = 2
	w = iota * 2 // w = 4
)

const x = iota // x = 0
const y = iota // y = 0

// 枚举 不像其他语言使用enum表示枚举，而是通过在const后跟一对圆括号定义一组常量的方式来实现枚举
// 如果赋值表达式和前一个一致，可以省略；
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)
