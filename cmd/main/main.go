package main

import (
	"fmt"
	"github.com/PhilLar/Golang_tasks/sqrt_gen"
	"github.com/PhilLar/Golang_tasks/sqrt_default"

)

func main() {
	fmt.Println("Sqrt as generator")
	for i := range sqrt_gen.Sqrt(2, 10) {
		fmt.Println(i)
	}

	fmt.Println("---------------\nSqrt as casual func")
	sqrt_default.Sqrt(2)
}