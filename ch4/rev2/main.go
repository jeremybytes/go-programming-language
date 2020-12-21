// Rev reverses an array of ints in place
package main

import "fmt"

func reverse(s *[7]int) {
	for i, j := 0, 6; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("Before: %v\n", a)
	reverse(&a)
	fmt.Printf("After: %v\n", a)
}
