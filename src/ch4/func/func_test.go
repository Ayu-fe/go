package func_test

import (
	"testing"
	"time"
)

func mutiReturnFunc() (int, int) {
	return 1, 2 // 多返回值
}

func addGetTiming(fun func(op int) int) func(op int) (int, float64) {
	return func(n int) (int, float64) {
		start := time.Now()
		result := fun(n)
		duration := time.Since(start).Seconds()
		return result, duration
	}
}

func slowFun(a int) int {
	time.Sleep(time.Second * 1)
	return a
}

func add(arr ...int) int { // 可变参数
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func TestMutiReturnFunc(t *testing.T) {
	// a, b := mutiReturnFunc()
	// t.Log(a, b)
	defer func() { // 最后执行 有点像final 即便是panic也会执行
		t.Log("over")
	}()
	newFunc := addGetTiming(slowFun)
	res, timing := newFunc(2)
	t.Log(res, timing)
	arrSum := add(1, 2, 3, 4)
	t.Log(arrSum)
}
