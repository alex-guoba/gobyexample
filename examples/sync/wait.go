package main

//import "fmt"
import "sync"
import "time"

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Wait()
		println("wait exit.")
	}()

	go func() {
		time.Sleep(time.Second)
		println("done.")
		wg.Done()
	}()

	wg.Wait()
	println("main exit.")
}
