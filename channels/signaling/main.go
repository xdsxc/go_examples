package main

import (
	"fmt"
	"time"
)

func workUntilClose(stop <-chan struct{}) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Printf("Hello, world!\n")
		case <-stop:
            fmt.Printf("received stop signal\n")
			return
		}
	}
}

func exampleChannelToClose() {
	const numGoRoutines = 3

	stop := make(chan struct{})
	for i := 0; i < numGoRoutines; i++ {
		go workUntilClose(stop)
	}

	time.Sleep(3 * time.Second)
	close(stop)
	time.Sleep(3 * time.Second)
}

func waitToStartWork(start <-chan struct{}) {
    fmt.Printf("Waiting for start signal\n")
	<-start
	fmt.Printf("Hello, world!\n")
}

func exampleChannelToStart() {
	const numGoRoutines = 3

	start := make(chan struct{})
	for i := 0; i < numGoRoutines; i++ {
		go waitToStartWork(start)
	}

	time.Sleep(3 * time.Second)
	close(start)
	time.Sleep(3 * time.Second)

}

func main() {
	exampleChannelToClose()
	exampleChannelToStart()
}
