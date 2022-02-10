package chapter

import "testing"

func TestAnimalMove(t *testing.T) {
	d := &Dog{
		feet:   4,
		Animal: &Animal{name: "MyDog"},
	}

	d2 := Dog{
		feet:   4,
		Animal: &Animal{name: "MyDog2"},
	}

	c := &Cat{
		feet:   4,
		Animal: &Animal{name: "MyCat"},
	}

	//c2 := Cat{
	//	feet:   4,
	//	Animal: &Animal{name: "MyCat2"},
	//}

	// Wolf 从 Animal 继承了 Mover 的方法实现
	w := &Wolf{&Animal{name: "MyWolf"}}

	AnimalMove(d)  // Dog MyDog moving...
	AnimalMove(d2) // Dog MyDog2 moving...
	AnimalMove(c)  // Cat MyCat moving...
	// AnimalMove(c2) // Type does not implement 'Mover' as the 'move' method has a pointer receiver
	AnimalMove(w) // Animal MyWolf moving...
}

func TestAnimalDemo(t *testing.T) {

	c := &Cat{
		feet:   4,
		Animal: &Animal{name: "MyCat"},
	}
	//d := &Dog{
	//	feet:   4,
	//	Animal: &Animal{name: "MyDog"},
	//}
	AnimalDemo(c)
	//AnimalDemo(d) // Type does not implement 'animal' as some methods are missing: say()
}

func TestEmptyInterface(t *testing.T) {
	EmptyInterface()
}

func TestShowEmpty(t *testing.T) {
	var y A
	y = "字符串2"
	ShowEmpty(y)
	y = 20
	ShowEmpty(y)
	y = false
	ShowEmpty(y)
}

func TestShowEmpty2(t *testing.T) {
	var y A
	y = "字符串2"
	ShowEmpty2(y) // string "字符串2"
	y = 20
	ShowEmpty2(y) // anything 20
	y = false
	ShowEmpty2(y) // anything false
}

func TestMapAny(t *testing.T) {
	MapAny()
}

func TestAssertType(t *testing.T) {
	var y A
	y = "字符串2"
	AssertType(y) // string 字符串2
	y = 20
	AssertType(y) // other ""
}
