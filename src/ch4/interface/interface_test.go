package interface_test

import (
	"fmt"
	"testing"
)

// 定义一个函数的类型
type Ifunc func(data int) int

// 定义一个接口
type IPeople interface {
	eat() string
}

// 定义struct
type People struct {
	name int
}

// struct中定义方法实现
func (e *People) eat() string {
	return "eat something"
}

func TestInterface(t *testing.T) {
	a := new(People)
	t.Log(a.eat())

	// 使用接口作为类型
	var b IPeople = new(People)
	// b.name = 12
	t.Log(b.eat())
	t.Log(b)
}

// 父类Animal
type Animal struct {
	name int
}

func (animal *Animal) speak() string {
	return "。。。。"
}

// 子类 Dog
type Dog struct {
	animal *Animal
}

func (dog *Dog) speak() string {
	// return "wang!"
	return dog.animal.speak()
}

// 子类 cat
type Cat struct {
	Animal // 这样写的话就实现了继承父类方法 可以自己重写活着复用
}

func (cat *Cat) speak() string {
	return "miao~"
}

func TestExtends(t *testing.T) {
	var animal *Animal = new(Animal)
	t.Log(animal.speak(), animal.name)

	cat := new(Cat)
	cat.name = 222
	t.Log(cat.speak(), cat.name)

	dog := new(Dog)
	t.Log(dog.speak())
}

// 练习多态
type Bird interface {
	speak() string
}

type maque struct{}

func (bird *maque) speak() string {
	return "zhizhizhi"
}

type yanzi struct{}

func (yanzi *yanzi) speak() string {
	return "xiuxiuxiu"
}

func TestDuotai(t *testing.T) {
	maque := new(maque)
	yanzi := new(yanzi)
	// t.Log(maque.speak())
	fmt.Printf("%T %v\n", maque, maque.speak())
	fmt.Printf("%T %v\n", yanzi, yanzi.speak())
	// t.Log(yanzi.speak())
}

// 练习断言 a.(type)
func judgeType(a interface{}) {

	// if i, ok := a.(int); ok {
	// 	fmt.Println("int", i)
	// 	return
	// }
	// if i, ok := a.(string); ok {
	// 	fmt.Println("string", i)
	// 	return
	// }
	// fmt.Println("unknow")
	switch v := a.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow")
	}
}
func TestJudgement(t *testing.T) {
	judgeType(10)
	judgeType("10")
	judgeType(true)
}
