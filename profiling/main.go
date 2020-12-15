package main

// https://golang.org/pkg/net/http/pprof/
import (
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func intAddr(v int) *int {
	return &v
}

func fibRecursive(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fibRecursive(n-1) + fibRecursive(n-2)
}

func workLoop1() {
	for range time.Tick(7 * time.Second) {
		fibRecursive(rand.Intn(100))
	}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		v := uint64(1)
	    for i := 0; i < 100000; i++ {
	        v *= 2
		}
	})

	go workLoop1()
	go func () {
		for range time.Tick(time.Millisecond) {
			http.Get("localhost:6060")
		}
	}()

	panic(http.ListenAndServe("localhost:6060", http.DefaultServeMux))
}



















































func workLoop2() {
	for range time.Tick(time.Millisecond) {
		fibMemoized(rand.Intn(100000))
	}
}

var mBuf []*int

func fibMemoized(n int) int {
	if mBuf == nil {
		mBuf = make([]*int, n+1)
		mBuf[0] = intAddr(0)
		mBuf[1] = intAddr(1)
	}

	if len(mBuf) <= n {
		buf := make([]*int, (n+1)*2)
		copy(buf, mBuf)
		mBuf = buf
	}

	if mBuf[n] != nil {
		return *mBuf[n]
	}

	mBuf[n] = intAddr(fibMemoized(n-1) + fibMemoized(n-2))
	return *mBuf[n]
}


