package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- work(j)
	}
}

func work(y int) int {
	return y * y
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
	
	for x := 1; x <= 10; x++ {
		jobs <- x
	}
	close(jobs)
	
	for a := 1; a <= 10; a++ {
		fmt.Println(<-results)
	}
}
