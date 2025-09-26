package main

import (
	"fmt"
	"sync"
	// "time"
)

// func worker(ch chan string) {
// 	ch <- "Task completed!"
// }

func Hi(i int) {
	fmt.Println("Hi", i)
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Hi(i)
		}(i)
	}

	wg.Wait()
	// time.Sleep(time.Second * 1)
	// messageChannel := make(chan string)
	// go worker(messageChannel)
	// msg := <-messageChannel
	// fmt.Println(msg)
}
