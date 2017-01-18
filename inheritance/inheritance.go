package main

import "fmt"

type One struct {
    Param11 int
    Param12 string
}

type Two struct {
    Param21 int
    Param22 string
}

type Three struct {
    One One
    Two Two
}

func main() {
    one := One{1,"1"}
    two := Two{2,"2"}
    three := Three{one, two}
    fmt.Println(three.One.Param11)
}
