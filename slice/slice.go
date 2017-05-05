package main

import "fmt"

func main() {
    test := []int{1, 2, 3}
    // Include 1 but exclude 3
    test1 := test[1:3]

    fmt.Println(test)
    fmt.Println(test1)
    test[1] = 0
    fmt.Println(test)
    fmt.Println(test1)

    test2 := make([]int, 5)
    fmt.Println(len(test2), cap(test2))
    test2 = append(test2, 1, 3)
    fmt.Println(len(test2), cap(test2))
}
