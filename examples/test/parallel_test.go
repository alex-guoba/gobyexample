package main

//import "fmt"
import "os"
import "time"
import "testing"

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
	t.Log("A quit")
}

func TestA1(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second * 2)
	t.Log("A1 quit")
}

func TestB(t *testing.T) {
	// go test -v -parallel 3 -args "b"
	if os.Args[len(os.Args)-1] == "b" {
		t.Parallel()
		t.Log("parallel in B")
	}

	time.Sleep(time.Second * 2)
	t.Log("B quit")
}

// test main
//
/*
func TestMain(m *testing.M) {
	tests := []testing.InternalTest{
		{"b", TestB},
		{"a", TestA},
	}

	benchmarks := []testing.InternalBenchmark{}
	examples := []testing.InternalExample{}
	match := func(pat, str string) (bool, error) {
		return true, nil
	}
	m = testing.MainStart(match, tests, benchmarks, examples)
	os.Exit(m.Run())
} */
