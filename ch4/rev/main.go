// Rev reverses a slice of ints in place
package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Before: %v\n", a)
	reverse(a[:])
	fmt.Printf("After: %v\n", a)
}
