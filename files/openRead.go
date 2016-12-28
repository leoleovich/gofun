package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, _ := ioutil.ReadFile("/tmp/test")
    lines := strings.Split(string(data), "\n")

    for _, line := range(lines) {
        fmt.Println(line)
    }
}
