package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var ops int64 = 0


	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	for r := 0; r < 10; r++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)

				mutex.Lock()
				state[key] = value
				mutex.Unlock()

				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println(atomic.LoadInt64(&ops))
	mutex.Lock()
	fmt.Println(state)
	mutex.Unlock()
}
