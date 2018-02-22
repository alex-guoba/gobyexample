package main

import "fmt"

func main() {
	aValue := new(int)

	// This will return 0, and not 100, as it is the default value for an integer.
	defer fmt.Println(*aValue)

	for i := 0; i < 100; i++ {
		*aValue++
	}
}
