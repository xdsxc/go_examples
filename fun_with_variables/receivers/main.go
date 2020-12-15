package main

import "fmt"

type Foo int
func (f Foo) Add(n int) {
    f += Foo(n)
}

func exampleBadIncrement() {
    var x Foo
    fmt.Printf("x = %d\n", x)
    x.Add(10)
    fmt.Printf("x = %d\n", x)
}

func main() {
    exampleBadIncrement()
}
