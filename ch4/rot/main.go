// Rot rotates a slice of ints
package main

import "fmt"

func rotateLeft(s []int, rotateBy int) {
	x := []int{}
	x = append(x, s[rotateBy:]...)
	x = append(x, s[:rotateBy]...)
	for i, n := range x {
		s[i] = n
	}
}

func rotateRight(s []int, rotateBy int) {
	r := len(s) - rotateBy
	rotateLeft(s, r)
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Before: %v\n", a)
	rotateRight(a[:], 3)
	fmt.Printf("After: %v\n", a)

	b := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Before: %v\n", b)
	rotateRight(b[:], 6)
	fmt.Printf("After: %v\n", b)
}
