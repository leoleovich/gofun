package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "bufio"
    "time"
)

func main() {
    file, _ := os.Open("/tmp/test")
    defer file.Close()

    // This one is faster a bit
    fmt.Fprintf(os.Stderr, time.Now().String()+"\n")
    data, _ := ioutil.ReadFile(file.Name())
    lines := strings.Split(string(data), "\n")

    for _, line := range(lines) {
        fmt.Println(line)
    }
    fmt.Fprintf(os.Stderr, time.Now().String()+"\n")

    s := bufio.NewScanner(file)
    for s.Scan() {
        fmt.Println(s.Text())
    }
    fmt.Fprintf(os.Stderr, time.Now().String()+"\n")
}
