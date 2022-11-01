package cap_test

import "testing"

func TestCap(t *testing.T) {
	// 切片声明  切片是容量 len是实际长度
	// 自增长 本质上是因为 用了多个连续的空间 然后把前一个的值复制过来去创建一个新的空间
	// 所以为什么是 a = append(a, 1)
	// 需要注意的是 这里的存储结构是共享的
	// 切片无法进行比较！！！
	a := []int{1, 2, 3, 4}
	t.Log(len(a), cap(a))

	b := []int{}
	for i := 0; i < 5; i ++ {
		b = append(b, i) // 类似于push, 切片专用, 数组无法自增长
	}
	t.Log(len(b), cap(b))


	c := make([]int, 5, 10) // 创建切片
	t.Log(len(c), cap(c))
	c = append(c, 1)
	t.Log(len(c), cap(c))

	d := [...]int{}
	// d = append(d, 1) 普通数组用不了
	t.Log(d, cap(d))
}