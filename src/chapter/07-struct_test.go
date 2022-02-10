package chapter

import (
	"fmt"
	"testing"
)

func TestAnonymous(t *testing.T) {
	Anonymous()
}

func TestInitialize2(t *testing.T) {
	Initialize2()
}

func TestInstance(t *testing.T) {
	Instance()
}

func TestPerson_Dream(t *testing.T) {
	p := NewPerson("test", "test", 18)
	p.Dream("something") // test dream: something
}

func TestPerson_SetAge(t *testing.T) {
	p := NewPerson("test", "test", 10)
	fmt.Printf("%#v\n", p) // &chapter.person{name:"test", city:"test", age:10}
	p.SetAge(18)
	fmt.Printf("%#v\n", p) // &chapter.person{name:"test", city:"test", age:18}
}

func TestMyInt_Say(t *testing.T) {
	var i MyInt = 10
	i.Say()

	j := MyInt(20)
	j.Say()
}

func TestDog_Wang(t *testing.T) {
	p := &Dog{
		feet:   4,
		Animal: &Animal{name: "MyDog"},
	}
	p.Wang()
	p.Move()
}
