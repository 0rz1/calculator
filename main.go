package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/0rz1/calculator/calc"
)

func removeNewLine(s string) string {
	for len(s) > 0 {
		c := s[len(s)-1]
		if c == '\r' || c == '\n' {
			s = s[:len(s)-1]
		} else {
			return s
		}
	}
	return s
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Calculator")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, rerr := reader.ReadString('\n')
		text = removeNewLine(text)
		if rerr != nil {
			fmt.Printf("%v\n", rerr)
		}
		result, err := calc.Calc(text)
		if err != nil {
			fmt.Printf("%v err: %v\n", text, err)
		} else {
			fmt.Printf("%v = %v\n", text, result)
		}
	}
}
