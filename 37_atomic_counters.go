package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	for i := 0; i < 50; i++ {
		// This will dispatch 50 workers each monotonically
		// incrementing the single uint64 counter reference
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
