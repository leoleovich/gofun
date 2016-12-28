package main
import (
    "fmt"
    "math/big"
)

func fact(factorial int, result *big.Int) {
    if factorial == 1 {
        return
    }
    result.Mul(big.NewInt(int64(factorial)),result)
    fact(factorial-1, result)
}

func main() {
    var factorial int
    result := big.NewInt(1)
    fmt.Scan(&factorial)
    fact(factorial, result)
    fmt.Print(result)
}
