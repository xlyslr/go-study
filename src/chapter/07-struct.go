package chapter

import "fmt"

// Anonymous 匿名结构体
func Anonymous() {
	// 如果折叠起来写，以;分隔
	var user struct {
		Name string
		Age  int
	}
	user.Name = "name"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

type person struct {
	name string
	city string
	age  int
}

func NewPerson(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// Instance 实例化
func Instance() {
	var p1 person
	p1.name = "1"
	fmt.Printf("%#v\n", p1) // chapter.person{name:"1", city:"", age:0}

	p2 := new(person)
	p2.name = "2"
	fmt.Printf("%#v\n", p2) // &chapter.person{name:"2", city:"", age:0}

	p3 := &person{}
	p3.name = "3"
	fmt.Printf("%#v\n", p3) // &chapter.person{name:"3", city:"", age:0}
}

// Initialize 结构体初始化
func Initialize2() {
	// 值结构体初始化
	p1 := person{
		name: "1",
		city: "2",
		age:  3,
	}
	fmt.Printf("%#v\n", p1) // chapter.person{name:"1", city:"2", age:3}
	// 引用类型结构体初始化

	p2 := &person{
		name: "1",
		city: "2",
		age:  3,
	}
	fmt.Printf("%#v\n", p2) // &chapter.person{name:"1", city:"2", age:3}

	p3 := NewPerson("1", "2", 3)
	fmt.Printf("%#v\n", p3) // &chapter.person{name:"1", city:"2", age:3}
}

func (p person) Dream(d string) {
	fmt.Printf("%s dream: %s\n", p.name, d)
}

func (p *person) SetAge(age int) {
	p.age = age
}

type MyInt int

func (i MyInt) Say() {
	fmt.Printf("%d\n", i)
}

type Animal struct {
	name string
}

func (a *Animal) Move() {
	fmt.Printf("%s 正在移动...\n", a.name)
}

type Dog struct {
	feet    int
	*Animal // 通过嵌套结构体来继承Animal的字段和方法
}

func (d *Dog) Wang() {
	fmt.Printf("%s 汪汪...\n", d.name)
}
