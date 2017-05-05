package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["QQQQ"] = 12

	fmt.Println(len(m))
}
