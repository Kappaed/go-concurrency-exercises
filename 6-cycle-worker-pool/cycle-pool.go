package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"
)

func main() {

	arr := []int{1,2,3,4,5,6,7,8,9,10}
	cycle := 0
	throttled := time.Tick(10*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mu sync.Mutex
	var wg sync.WaitGroup
	go func(throttler <-chan time.Time) {
		for range throttler {
			mu.Lock()	
			for i := range arr {
				arr[i] += 1
			}
			cancel()
			mu.Unlock()
		}
		
	}(throttled)
	for {
		for i,_ := range arr {
			wg.Add(1)
			go func(idx int) {
				defer wg.Done()
				select {
					case <- ctx.Done():
						log.Println("cancelled.")
						return
					default:
						log.Println(arr[idx])				
				}
			}(i)
			
		}
		wg.Wait()
		ctx, cancel = context.WithCancel(context.Background())
		log.Println("cycled updated:", cycle)
		fmt.Printf("num of goroutines currently running: %d\n",runtime.NumGoroutine())
		time.Sleep(3*time.Second)
		cycle += 1

	}

}