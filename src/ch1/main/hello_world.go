package main

import (
	"fmt"
	"os"
)

func main() {
	
	if len(os.Args) > 1 {
		fmt.Println("Hello, World", os.Args[1])
	}
	// os.Exit(-1) // 返回值
}