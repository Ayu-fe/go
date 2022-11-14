package error_test

import (
	"errors"
	"testing"
)

func getFiber(n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("小于0")
	}
	fiber := []int{1, 1}
	for i := 2; i < n; i++ {
		fiber = append(fiber, fiber[i-1]+fiber[i-2])
	}
	return fiber, nil
}

func TestFiber(t *testing.T) {
	defer func() {
		t.Log("我在执行了")
	}()
	if list, err := getFiber(1); err != nil {
		t.Log(err)
	} else {
		t.Log("niubi", list)
	}
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Log("捕获错误") // 看起来是try catch
		}
	}()
	panic(errors.New("报错啦")) // 可捕获的错误
}
