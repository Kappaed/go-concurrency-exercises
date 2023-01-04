package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func worker(jobs <-chan string, result chan<- string) {
		for {
			_, ok := <- jobs
			if(!ok) {
				return
			}
			result <- "done"
		}
}

const jobs = 10
const finishAfter = 5*time.Second

func main() {
	j := make(chan string,jobs)
	r := make(chan string,jobs)
	for i :=0;i<jobs;i++ {
		go worker(j,r)
	}

	for i:=0;i<jobs;i++ {
		j <- "job"
	}
	close(j)
	
	finishTicker := time.NewTicker(finishAfter)
	for {
		select {
			case job,ok := <- r:
				if (!ok) {
					break
				} 
				finishTicker.Reset(finishAfter)
				log.Println(job)
			case <- finishTicker.C:
				close(r)
				log.Println("finished")
				fmt.Printf("num of goroutines currently running: %d\n",runtime.NumGoroutine())
				return
		}
	}


}