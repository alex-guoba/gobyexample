// +build release

package main

import "fmt"

// use "go build -tags \"release\"" to build release version
func hello() {
	fmt.Println("release version.")
}
