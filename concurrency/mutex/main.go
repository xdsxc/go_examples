package main

import (
    "fmt"
    "sync"
)

var (
	value int
)

func waitToStartWork(start <-chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    <-start
    value += 1
}

func exampleDeadlock() {
    const (
        numGoRoutines  = 1000
    )

    start := make(chan struct{})
    wg := sync.WaitGroup{}
    wg.Add(numGoRoutines)
    for i := 0; i < numGoRoutines; i++ {
        go waitToStartWork(start, &wg)
    }
    close(start)
    wg.Wait()
    fmt.Printf("value = %d\n", value)
}

func main() {
    exampleDeadlock()
}
