package main

import "fmt"

type tdata struct {
	x int
}

func main() {
	d := tdata{100}

	var t interface{} = d // d{ will be copyed when assigned to interface var

	fmt.Println(t.(tdata).x) // t.data unaddressable

	var nt interface{} = &d
	nt.(*tdata).x = 200
	fmt.Println(nt.(*tdata).x)
}
