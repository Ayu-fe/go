package for_test

import "testing"

func TestForIf(t *testing.T) {
	// 循环 for
	for i := 0; i < 5; i ++ {
		t.Log(i)
	}
	n := 0
	for n < 5 {
		t.Log(n)
		n ++;
	}

	// 条件
	if a := 1 == 1; a {
		t.Log(a)
	}
	if 1 == 1 {
		t.Log("ok")
	}
	// 一般用在多结果返回
	// if b, err := func(); err == nil {
	// 	t.Log(b)
	// }

	// switch
	for p := 0; p < 5; p ++ {
		// switch p {
		// case 1, 2: // case支持多个
		// 	t.Log("牛逼")
		// case 3, 4:
		// 	t.Log("xxx")
		// default:
		// 	t.Log("default")
		// }

		// 还可以这么写 感觉像if else
		switch {
		case p == 1 || p == 2:
			t.Log("牛逼")
		case p == 3:
			t.Log("离哈")
		}

		// 还可以这么写 非常方便
		// switch os := runtime.GOOS; os {
		// case "darwin":
		// 	t.Log("darwin")
		// default:
		// 	t.Log("default")
		// }

	}
}