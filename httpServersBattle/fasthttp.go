package main

import (
	"github.com/valyala/fasthttp"
	"fmt"
)

func handleRequest(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "Hello, World!")
}

func main() {
	fmt.Println(fasthttp.ListenAndServe(":7000", handleRequest))
}