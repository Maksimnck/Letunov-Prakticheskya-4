package main

import (
	"fmt"
	"time"
)

type source struct {
	name  string
	delay time.Duration
	ok    bool
}

type result struct {
	source string
	data   string
	ok     bool
}

func search(s source, query string) result {
	time.Sleep(s.delay)
	if !s.ok {
		return result{source: s.name, ok: false}
	}
	return result{source: s.name, data: "найдено: " + query, ok: true}
}

func firstResult(query string) (result, bool) {
	sources := []source{
		{name: "база данных 1", delay: 500 * time.Millisecond, ok: false},
		{name: "апи 1", delay: 800 * time.Millisecond, ok: true},
		{name: "база данных 2", delay: 300 * time.Millisecond, ok: false},
		{name: "апи 2", delay: 1200 * time.Millisecond, ok: true},
	}

	results := make(chan result, len(sources))
	for _, s := range sources {
		go func(s source) {
			results <- search(s, query)
		}(s)
	}

	timeout := time.After(3 * time.Second)
	for range sources {
		select {
		case r := <-results:
			if r.ok {
				return r, true
			}
			fmt.Println(r.source, "-не найдено ничего")
		case <-timeout:
			return result{}, false
		}
	}
	return result{}, false
}

func main() {
	r, ok := firstResult("golang")
	if !ok {
		fmt.Println("нет результата")
		return
	}
	fmt.Printf("первый результат от %s: %s\n", r.source, r.data)
}
