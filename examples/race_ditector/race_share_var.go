package main

import "os"
import "fmt"

// ParallelWrite writes data to file1 and file2, returns the errors.
func ParallelWrite(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("file1")
	fmt.Printf("f1 err %x\n", &err)
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			fmt.Printf("f1 in go err %x\n", &err)
			res <- err
			f1.Close()
		}()
	}
	f2, err := os.Create("file2") // The second conflicting write to err.
	fmt.Printf("f2 err %x\n", &err)
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			fmt.Printf("f2 in go err %x\n", &err)
			res <- err
			f2.Close()
		}()
	}
	return res
}

func main() {

	// go run -race ./race_map.go

	data := make([]byte, 100)
	copy(data, "hello")
	cerr := ParallelWrite(data)
	<-cerr
	<-cerr
}
