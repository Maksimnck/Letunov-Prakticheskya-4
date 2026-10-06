package main

import "fmt"
import "time"

func main() {
	requests := make(chan int, 15)
	for i := 1; i<=15;i++ {
		requests <- i
	}
	close(requests)
	tick := time.Tick(200 * time.Millisecond)
	for req := range requests {
		<-tick
		fmt.Println("обработан запрос", req)
	}
}
