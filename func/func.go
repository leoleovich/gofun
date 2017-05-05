package main

import "fmt"

func main() {
	testVar := 2
	testFunc := func(num1, num2 int) int {
		defer func() {
			recover()
		}()
		fmt.Println(testVar)
		return num1/num2
	}

	fmt.Println(testFunc(1,0))
}
