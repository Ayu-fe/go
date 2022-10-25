package data

import "testing"

type MyInt int64

func TestData(t *testing.T) {
	// 无法隐式类型转换
	// 即便是别名也不行
	var a int64 = 1
	var b MyInt
	b = MyInt(a) // 无法转换
	t.Log(a, b)
}

func TestPoint(t *testing.T) {
	a := 1
	aaddr := &a
	// go语言无法使用指针进行运算
	t.Log(a, aaddr)
	t.Logf("%T %T", a, aaddr)
}