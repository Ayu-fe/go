package struct_test

import "testing"

type People struct {
	name string
	age  int
	dead bool
}

// 给struct上定义一个函数
func (people *People) eat() string { // 这样创建的使用的是相同的地址
	people.age = 123
	return "people"
}
func (people People) eat2() string { // 这样创建会复制一份 通常使用上面的定义方式
	people.age = 123
	return "people"
}

func TestStruct(t *testing.T) {
	// 几种不同的初始化
	people := People{"aaa", 12, true}
	t.Log(people)
	people2 := People{name: "lism", age: 18}
	t.Log(people2, people2.name)

	people3 := new(People) // 创建一个指向People实例的指针
	people3.name = "wfq"

	t.Log(people2.eat())
	// t.Log(people3)
}
