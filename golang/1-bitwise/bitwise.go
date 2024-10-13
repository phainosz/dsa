package main

import "fmt"

func main() {
	// a = 5(00000101), b = 9(00001001)
	a, b := 5, 9

	fmt.Printf("a value: %d\n", a)
	fmt.Printf("a in binary: %08b\n", a)
	fmt.Printf("a in binary left shifting: %08b\n", a<<1)
	fmt.Printf("a value after left shifting: %d\n\n", a<<1)

	fmt.Printf("b value: %d\n", b)
	fmt.Printf("b in binary: %08b\n", b)
	fmt.Printf("b in binary left shifting: %08b\n", b<<1)
	fmt.Printf("b value after left shifting: %d\n", b<<1)

	//outputs
	//a value: 5
	//a in binary: 00000101
	//a in binary left shifting: 00001010
	//a value after left shifting: 10

	//b value: 9
	//b in binary: 00001001
	//b in binary left shifting: 00010010
	//b value after left shifting: 18
}
