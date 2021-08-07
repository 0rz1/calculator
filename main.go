package main

import (
	"fmt"

	"github.com/0rz1/calculator/calc"
)

func main() {
	res, err := calc.Calc("1 + 2")
	fmt.Println(res, err)
}
