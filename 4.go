package main

import "fmt"
import "net/http"
import "sync"


type item struct {
	url    string
	status string
	err    error
}

func worker(urls <-chan string, results chan<- item, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		var it item
		it.url = url
		resp, err := http.Get(url)
		if err != nil {
			it.err = err
			results <- it
			continue
		}
		it.status = resp.Status
		resp.Body.Close()
		results <- it
	}
}

func main() {
	list := []string{
		"https://go.dev",
		"https://github.com",
		"https://google.com",
		"https://yandex.ru",
	}

	urls := make(chan string, len(list))
	results := make(chan item, len(list))
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(urls, results, &wg)
	}

	for _, u := range list {
		urls <- u
	}
	close(urls)

	go func() {
		wg.Wait()
		close(results)
	}()

	for it := range results {
		if it.err != nil {
			fmt.Printf("%s\tошибка: %v\n", it.url, it.err)
			continue
		}
		fmt.Printf("%s\t%s\n", it.url, it.status)
	}
}
