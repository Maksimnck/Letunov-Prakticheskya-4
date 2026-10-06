package main

import "fmt"
import "sync"
import "time"

func main() {
	var wg sync.WaitGroup 
    wg.Add(1)
	go func() {
		for x := 1; x <= 5; x++ {
			fmt.Println(x);
			time.Sleep(time.Second)
		}
		wg.Done()
	}()
	
	wg.Wait()
}
