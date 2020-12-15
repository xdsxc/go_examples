package main

import (
    "fmt"
    "sync"
)

func exampleClosureBug() {
    const (
        numGoRoutines = 10
    )

    wg := sync.WaitGroup{}
    wg.Add(numGoRoutines)
    start := make(chan struct{})
    for i := 0; i < numGoRoutines; i++ {
        go func() {
            <-start
            fmt.Printf("i = %d\n", i)
            wg.Done()
        }()
    }

    close(start)
    wg.Wait()
}

func main() {
    exampleClosureBug()
}
