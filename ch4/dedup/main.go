// Dedup eliminates adjacent duplicates in a []string slice
package main

import "fmt"

func dedup(strings []string) (out []string) {
	out = strings[:0]
	prev := ""
	for _, s := range strings {
		if s != prev {
			out = append(out, s)
		}
		prev = s
	}
	return
}

func main() {
	a := []string{"Hello", "Hello", "World", "World"}
	fmt.Printf("Before: %v\n", a)
	a = dedup(a)
	fmt.Printf("After:  %v\n", a)
}
