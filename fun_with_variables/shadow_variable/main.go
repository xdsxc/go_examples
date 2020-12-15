package main

import "fmt"

func exampleShadowVariable() {
    n := 5
    addN := func(n int) {
        n += 100
    }

    addN(n)
    fmt.Printf("n = %d\n", n)

    n = 10
    addN(n)
    fmt.Printf("n = %d\n", n)
}

func main() {
    exampleShadowVariable()
}
