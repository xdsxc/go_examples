package buffered

import (
	"fmt"
	"time"
)

func exampleChannelConcurrencyLimiter() {
	const (
		maxConcurrency = 3
		numGoRoutines  = 10
	)

	limiter := make(chan struct{}, maxConcurrency)
	for i := 0; i < numGoRoutines; i++ {
		go func() {
			for {
				limiter<- struct{}{}
				fmt.Printf("Work in progress...\n")
				time.Sleep(2*time.Second)
				fmt.Printf("Done\n")
				<-limiter
			}
		}()
	}

	time.Sleep(10*time.Second)
}

func main() {
	exampleChannelConcurrencyLimiter()
}
