package main

import "fmt"

func main() {
    test := []int{1,2,3}
    test1 := test

    fmt.Println(test)
    fmt.Println(test1)
    test[1] = 0
    fmt.Println(test)
    fmt.Println(test1)
}
