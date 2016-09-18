package main

import "testing"
import "time"

func add(x, y int) int {
	return x + y
}

func BenchmarkAdd(b *testing.B) {
	time.Sleep(time.Second)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = add(1, 2)
	}
}
