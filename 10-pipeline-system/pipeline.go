package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup
	jobs := make(chan int, 10)
	results := make(chan int, 10)
	for i:=0;i<10;i++ {
		wg.Add(2)
		go func(jchan <-chan int, rchan chan<- int) {
			defer wg.Done()
			job, ok := <-jchan
			if (!ok) {
				return
			} 
			time.Sleep(1*time.Second)
			rchan <- job*job
			
		}(jobs, results)

		go func(rchan <- chan int) {
			defer wg.Done()
			job, ok := <- rchan
			if (!ok) {
				return
			}
			fmt.Println(job*job)

		}(results)
	}
	for i:=0;i<10;i++ {
		jobs <- i+1
	}
	close(jobs)
	wg.Wait()
	close(results)

	fmt.Printf("num of goroutines currently running: %d\n",runtime.NumGoroutine())
}