package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
)

type Being struct {
    Name   string
    Age    int
    Weight int
    Next * Being
}

func reverse(curr * Being) *Being {
    if curr.Next == nil {
        return curr
    } else {
        newHead := reverse(curr.Next)
        curr.Next.Next = curr
        curr.Next = nil
        return newHead
    }
}

func main() {

    b := Being{}
    data, _ := ioutil.ReadFile("./test.txt")
    for _, line := range strings.Split(string(data), "\n") {
        dataSplit := strings.Split(line, " ")
        age, _ := strconv.Atoi(dataSplit[1])
        weight, _ := strconv.Atoi(dataSplit[2])
        bng := Being{dataSplit[0], age, weight, b.Next}
        b.Next = &bng
    }

    bRev := reverse(&b)
    for s := bRev; s != nil; s = s.Next {
        fmt.Println(s.Name)
    }
}