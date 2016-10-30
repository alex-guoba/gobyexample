// +build log

package main

import "fmt"

// go build -tags "release log"
func init() {
	fmt.Println("logging")
}
