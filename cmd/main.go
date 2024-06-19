package main

import (
	"GoHexArch/internal/adaptors/core/arithmetic"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	arithmeticAdapter := arithmetic.NewAdapter()
	result, err := arithmeticAdapter.Div(10, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
