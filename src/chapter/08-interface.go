package chapter

import "fmt"

type Mover interface {
	// 不需要 func，因为都是func
	move()
}

func (d Dog) move() {
	fmt.Printf("Dog %s moving...\n", d.name)
}

type Cat struct {
	feet int
	*Animal
}

func (c *Cat) move() {
	fmt.Printf("Cat %s moving...\n", c.name)
}

type Wolf struct {
	*Animal
}

func (a *Animal) move() {
	fmt.Printf("Animal %s moving...\n", a.name)
}

func AnimalMove(mover Mover) {
	mover.move()
}

type Sayer interface {
	say()
}

type animal interface {
	Sayer
	Mover
}

func (c *Cat) say() {
	fmt.Printf("Cat %s miao...\n", c.name)
}

func AnimalDemo(a animal) {
	a.say()
	a.move()
}

// A 这是一个空接口
type A interface {
}

// B 这是一个空结构体
type B struct {
}

func EmptyInterface() {
	// 这是一个匿名空接口变量
	var x interface{}
	x = "字符串1"
	fmt.Printf("%#v\n", x)
	x = 10
	fmt.Printf("%#v\n", x)
	x = true
	fmt.Printf("%#v\n", x)

}

func AssertType(a A) {
	// s是断言之后的值，如果断言失败则为空；如果断言成功ok为true，否则为false
	s, ok := a.(string)
	if ok {
		fmt.Printf("string %s\n", s)
	} else {
		fmt.Printf("other %#v\n", s)
	}
}

func ShowEmpty(a A) {
	fmt.Printf("%#v\n", a)
}

func ShowEmpty2(a A) {
	switch a.(type) {
	case string:
		fmt.Printf("string %#v\n", a)
	default:
		fmt.Printf("anything %#v\n", a)
	}
}

func MapAny() {
	m := make(map[string]interface{})
	m["a"] = "字符串"
	m["b"] = 10
	m["c"] = true
	for s, i := range m {
		fmt.Printf("%s = %#v\n", s, i)
	}
}
