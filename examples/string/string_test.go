package main

import "testing"

func test() string {
	var s string
	for i := 0; i < 1000; i++ {
		s += "a"
	}
	return s
}

func BenchmarkTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		test()
	}
}
