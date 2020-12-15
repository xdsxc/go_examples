package main

import (
    "fmt"
    "time"
)

func waitToStartWork(start <-chan struct{}) {
    fmt.Printf("Waiting for start signal\n")
    <-start
    time.Sleep(2*time.Second)
    fmt.Printf("Hello, world!\n")
}

func exampleWaitGroup() {
    const (
        numGoRoutines  = 3
    )

    start := make(chan struct{})
    for i := 0; i < numGoRoutines; i++ {
        go waitToStartWork(start)
    }

    close(start)
    time.Sleep(1*time.Second)
}

func main() {
    exampleWaitGroup()
}
