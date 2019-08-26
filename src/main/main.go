package main

import (
	"fmt"
	"sqrt_gen"
	"sqrt_default"

)

func main() {
	fmt.Println("Sqrt as generator")
	for i := range sqrt_gen.Sqrt(2, 10) {
		fmt.Println(i)
	}

	fmt.Println("---------------\nSqrt as casual func")
	sqrt_default.Sqrt(2)
}