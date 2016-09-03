package main

import "fmt"
import "reflect"
import "unsafe"

func pp(format string, ptr interface{}) {
	p := reflect.ValueOf(ptr).Pointer()
	h := (*uintptr)(unsafe.Pointer(p))
	fmt.Printf(format, *h)
}

func toString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

func main() {
	s := "hello"
	pp("s: %x\n", &s)

	// type conversation will lead to mem alloc & value copy
	bs := []byte(s)
	s2 := string(bs)
	pp("string to []byte, bs: %x\n", &bs)
	pp("string to string, s2: %x\n", &s2)

	// []byte and string have the same header(particial), it's not safe, but effeciency
	bs_new := []byte("hello again")
	s_new := toString(bs_new)
	pp("bs_new address: %x\n", &bs_new)
	pp("s_new address: %x\n", &s_new)
}
