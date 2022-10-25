package test

import (
	"testing" 
	"fmt"
)

func TestFirstTry(t *testing.T) {
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
