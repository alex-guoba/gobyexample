// Go supports [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function),
// which can form <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>.
// Anonymous functions are useful when you want to define
// a function inline without having to name it.

package main

import "fmt"

// This function `intSeq` returns another function, which
// we define anonymously in the body of `intSeq`. The
// returned function _closes over_ the variable `i` to
// form a closure.
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func test() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		s = append(s, func() {
			// read the last i when executed
			fmt.Println(&i, i)
		})
	}
	return s
}

func test2() []func() {
	var s []func()

	for i := 0; i < 2; i++ {
		x := 1 // differnt closure env now
		s = append(s, func() {
			fmt.Println(&x, x)
		})
	}
	return s
}

func test3(x int) (func(), func()) {
	return func() {
			fmt.Println(x)
			x += 11
		},
		func() {
			fmt.Println(x)
		}
}

func main() {

	// We call `intSeq`, assigning the result (a function)
	// to `nextInt`. This function value captures its
	// own `i` value, which will be updated each time
	// we call `nextInt`.
	nextInt := intSeq()

	// See the effect of the closure by calling `nextInt`
	// a few times.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// To confirm that the state is unique to that
	// particular function, create and test a new one.
	newInts := intSeq()
	fmt.Println(newInts())

	for _, f := range test() {
		f()
	}

	for _, f := range test2() {
		f()
	}

	a, b := test3(10)
	a()
	b()
}
