package try_test

import (
	"testing" 
	"fmt"
)

const (
	mon = iota
	tue
	wes
)

// iota 0 1 2 ...
const (
	Readable = 1 << iota // 0001
	Writeable						 // 0010
	Excuteable           // 0100
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
	c := 7 // 0111
	t.Log(c & Readable, c & Writeable, c & Excuteable) // 1 2 4
	d := 1 // 0001
	t.Log(d & Readable, d & Writeable, d & Excuteable) // 1 0 0
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
