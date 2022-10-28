package array_test

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int// 初始化数组默认为0
	var arrStr [5]string

	tArr := [3]int{1, 4, 4} // 初始化并且赋值
	martixArr := [3][3]int{{1, 3, 2}, {5, 3, 6}}

	for i := 0; i < len(arr); i ++ {
		t.Log(arr[i])
	}
	for i := 0; i < len(arrStr); i ++ {
		t.Log(arrStr[i])
	}
	for i := 0; i < len(tArr); i ++ {
		t.Log(tArr[i])
	}
	for i := 0; i < len(martixArr); i ++ {
		t.Log(martixArr[i])
	}

	arrb := [...]int{9, 9, 9} // 不知道具体个数的声明

	for idx, e := range arrb { // 类似foreach的写法
		t.Log(idx, e)
	}
	for _, e := range arrb { // 使用占位
		t.Log(e)
	}
	arrc := arrb[0:2] // 数组的截取 左闭右开区间 类似于slice 不改变原数组 不过居然不是创建一个新的 而是使用了原数组的引用
	arrd := arrb[:] // 可以简写 表示开始结束
	arrb[0] = 4
	arrd[1] = 7
	t.Log(arrc, arrb, arrd)
	// [4 7] [4 7 9] [4 7 9]
}

