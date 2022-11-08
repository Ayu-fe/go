package map_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	obj := map[string]int{"a": 123, "b": 666}

	t.Log(obj, obj["a"])

	// map中不存在的值初始化为0
	obj2 := make(map[int]int, 10)
	obj2[1] = 1
	t.Log(obj2, len(obj2), obj2[2])

	if v, ok := obj2[1]; ok { // 手动判断一下 是否存在这个值，有点牛逼啊
		t.Log(v)
		t.Logf("值为 %d", v) // 这个让我想起c语言
	} else {
		t.Log("unknow")
	}

	// 使用rang遍历map
	for key, val := range obj {
		t.Log(key, val)
	}

	obj5 := map[int]func(data int) int{} // map中放函数
	obj5[1] = func(data int) int {
		return data
	}
	obj5[2] = func(data int) int {
		return data * 2
	}
	t.Log(obj5[1](2), obj5[2](2))

	obj6 := map[int]bool{}
	obj6[1] = true
	delete(obj6, 1) // 删除map中的key

}

func TestString(t *testing.T) {
	str := ""
	str = "中,国"          // 字符串本质上存在slice里 存的是byte数
	t.Log(str, len(str)) // len求出来的是byte数
	// arr := []string{str}
	// arr2 := []rune(str) // 得到str的unicode

	// t.Log(arr, len(arr), arr[0], arr2)
	// t.Logf("%x", str) // 输出utf8
	// t.Log(len(str))
	// strings strconv

	for _, i := range str {
		t.Log(i)
	}

	temp := strings.Split(str, ",") // split
	t.Log(temp)
	for _, i := range temp {
		t.Log(i)
	}
	temp2 := strings.Join(temp, "-") // join
	t.Log(temp2)

	strNum := "牛逼"
	if val, err := strconv.Atoi(strNum); err == nil { // 字符串转整形 注意返回值
		t.Log(val)
	} else {
		t.Log(err)
	}

	num := 10
	bb := strconv.Itoa(num) // 整形转字符串
	t.Log(bb + "11")

}
