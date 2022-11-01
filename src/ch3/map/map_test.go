package map_test

import "testing"

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
}