package struct_test

import "testing"

type People struct {
	name string
	age  int
	dead bool
}

func (people People) String() string {
	return
}

func TestStruct(t *testing.T) {
	// 几种不同的初始化
	people := People{"aaa", 12, true}
	t.Log(people)
	people2 := People{name: "lism", age: 18}
	t.Log(people2, people2.name)

	people3 := new(People) // 创建一个只想People实例的指针
	people3.name = "wfq"
	t.Log(people3)

}
