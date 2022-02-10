package chapter

import "fmt"

func main() {
	IntType()
	BoolType()
	StringType()
}
func IntType() {
	// 根据编译目标系统决定长度，32位系统是int32，,64位系统是int64
	var a int
	a = 1234
	// 取值-128 ~ 127
	var b int8
	b = -128
	// Java中的short
	var c int16
	// Java中的int
	var d int32
	// Java中的long
	var e int64
	fmt.Printf("%d %d %d %d %d\n", a, b, c, d, e)
}

func BoolType() {
	var a, b = true, false
	fmt.Printf("%t %t\n", a, b)
}

func StringType() {
	var a = "测试"
	var b = `测试
			多行
			文本`
	fmt.Printf("%s %s\n", a, b)
	c := a + " " + b
	fmt.Printf("%s\n", c)
}

func DataTrans() {
	a := 1
	b := float64(a)
	fmt.Printf("%d\n", int32(b))
}
