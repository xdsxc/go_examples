package exdeadlock

import (
    "fmt"
    "sync"
)

var (
	exampleMutexValue int
	mtx1 sync.Mutex
	mtx2 sync.Mutex
)

func waitToStartWork1(start <-chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    <-start
    mtx1.Lock()
    mtx2.Lock()
    exampleMutexValue += 1
    mtx1.Unlock()
    mtx2.Unlock()
}

func waitToStartWork2(start <-chan struct{}, wg *sync.WaitGroup) {
    defer wg.Done()
    <-start
    mtx2.Lock()
    mtx1.Lock()
    exampleMutexValue += 1
    mtx2.Unlock()
    mtx1.Unlock()
}

func main() {
    const (
        numGoRoutines  = 10000
    )

    start := make(chan struct{})
    wg := sync.WaitGroup{}
    wg.Add(numGoRoutines)
    for i := 0; i < numGoRoutines; i++ {
        if i % 2 == 0 {
            go waitToStartWork1(start, &wg)
        } else {
            go waitToStartWork2(start, &wg)
        }
    }
    close(start)
    wg.Wait()
    fmt.Printf("exampleMutexValue = %d\n", exampleMutexValue)
}
