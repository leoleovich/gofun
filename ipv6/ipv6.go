package main

import (
    "os"
    "fmt"
    "strings"
)

func main() {
    args := os.Args[1:]
    ipv6delimiter := ":"

    if len(args) != 1 {
        fmt.Println("Specify only 1 param please")
        os.Exit(1)
    }
    ipv6str := strings.Replace(args[0], ipv6delimiter+ipv6delimiter, ":xxxx:", -1)
    ipv6raw := strings.Split(ipv6str,ipv6delimiter)
    var ipv6full []string

    for i, part := range(ipv6raw) {
        if part == "xxxx" {
            /*
            8 (len of ipv6)
            -i(How many were before)
            -(len(origin string)-How many were before, to get how many will be after
            +1 for current position
            */
            for n:=0; n<8-i-(len(ipv6raw)-i)+1 ; n++ {
                ipv6full = append(ipv6full, "0000")
            }
        } else if part < "ffff"{
            for k:=len(part); k<4; k++ {
                part = "0" + part
            }
            ipv6full = append(ipv6full, part)
        } else {
            fmt.Println("ipv6 can not exceed ffff for 1 group: " + part)
            os.Exit(1)
        }
    }

    fmt.Println(strings.Join(ipv6full, ":"))
}
