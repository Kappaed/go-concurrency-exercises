//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, msgChannel chan<- *Tweet) {
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			fmt.Println("closing channel...")
			close(msgChannel)
			return
		}
		msgChannel <- tweet
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	messages := make(chan *Tweet)
	var wg sync.WaitGroup

	for i:=0;i<5;i++ {
		wg.Add(1)
		go func(msgChannel <-chan *Tweet) {
			for {
				select {
					case tweet, open := <-msgChannel:
						if !open {
							wg.Done()
							return
						}
						if tweet.IsTalkingAboutGo() {
							fmt.Println(tweet.Username, "\ttweets about golang")
						} else {
							fmt.Println(tweet.Username, "\tdoes not tweet about golang")
						}
				}
			}				
		}(messages)

	}
	// Producer
	producer(stream, messages)
	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
