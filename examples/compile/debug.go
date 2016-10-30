// +build !release

package main

import "fmt"

// use "go build" to build debug version
func hello() {
	fmt.Println("debug version.")
}
