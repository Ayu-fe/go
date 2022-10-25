package try_test

import (
	"testing" 
	"fmt"
)

const (
	mon = iota + 1
	tue
	wes
)

func TestFirstTry(t *testing.T) {
	// var cc int = 1
	a := 1
	t.Log(a)
	b := 1
	fmt.Println(a)
	for i := 0; i < 5; i ++ {
		fmt.Println(" ", b)
		temp := a
		a = b
		b = temp + a
	}
}

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	// temp := a
	// a = b 
	// b = temp
	a, b = b, a
	t.Log(a, b)
}

func TestConst(t *testing.T) {
	t.Log(mon, tue, wes)
}
