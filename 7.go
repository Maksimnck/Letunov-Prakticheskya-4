package main

import "fmt"
import "sync"

type op struct {
	kind string
	key  string
	val  int
	resp chan int
}

func manager(ops <-chan op) {
	state := make(map[string]int)
	for o := range ops {
		switch o.kind {
		case "get":
			o.resp <- state[o.key]
		case "set":
			state[o.key] = o.val
			o.resp <- state[o.key]
		case "add":
			state[o.key] += o.val
			o.resp <- state[o.key]
		}
	}
}

func get(ops chan<- op, key string) int {
	resp := make(chan int)
	ops <- op{kind: "get", key: key, resp: resp}
	return <-resp
}

func set(ops chan<- op, key string, val int) int {
	resp := make(chan int)
	ops <- op{kind: "set", key: key, val: val, resp: resp}
	return <-resp
}

func add(ops chan<- op, key string, val int) int {
	resp := make(chan int)
	ops <- op{kind: "add", key: key, val: val, resp: resp}
	return <-resp
}

func main() {
	ops := make(chan op)
	go manager(ops)

	set(ops, "counter", 0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				add(ops, "counter", 1)
			}
		}()
	}
	wg.Wait()

	fmt.Println("counter =", get(ops, "counter"))
	close(ops)
}
