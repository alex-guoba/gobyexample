package main

import "time"

func main() {
	m := make(map[string]int)

	// Maps are not safe for concurrent use: it's not defined what happens when you read and write to them simultaneously. If you need to read from and write to a map from concurrently executing goroutines, the accesses must be mediated by some kind of synchronization mechanism. One common way to protect maps is with sync.RWMutex.

	// go run -race ./race_map.go
	go func() {
		for {
			m["a"] += 1
			time.Sleep(time.Microsecond)
		}
	}()

	go func() {
		for {
			_ = m["b"]
			time.Sleep(time.Microsecond)
		}
	}()

	select {}
}
