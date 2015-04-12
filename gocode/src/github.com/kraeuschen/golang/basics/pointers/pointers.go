package main

import "fmt"

// copy of ival distinct
func zeroval(ival int) {
	ival = 0
}

// code dereferences the pointer from its memory address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // &i syntax gives the memory address of i
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)
}
