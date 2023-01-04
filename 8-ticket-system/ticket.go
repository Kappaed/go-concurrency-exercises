package main

import (
	"fmt"
	"runtime"
	"sync"
)



const maxTicketNum = 10
func main() {
	var ticketMap = make(map[int]int)
	var wg sync.WaitGroup
	var mu sync.Mutex
	tickets := []int{3,3,3,3,2,2,2,2,1,1,1,1,1,1,4,4,4,4,6,6,6,6,7,7,7,7,9,9,9,9,9,10,10,1,1,1,6,6,6,8}
	for i:=1;i<=maxTicketNum;i++ {
		ticketMap[i] = 0
	}

	for _,val := range tickets {

		wg.Add(1)
		go func(ticketNum int) {
			mu.Lock()
			defer mu.Unlock()
			defer wg.Done()
			count, ok := ticketMap[ticketNum]
			if(!ok) {
				return
			}

			ticketMap[ticketNum] =  count+1
		}(val)
	}

	wg.Wait()
	for i:=1;i<=maxTicketNum;i++ {
		val, ok := ticketMap[i]
		if (!ok) {
			continue
		}
		fmt.Printf("ticket num %d: %d\n", i, val)
	}
	
	fmt.Printf("num of goroutines currently running: %d\n",runtime.NumGoroutine())

	

}