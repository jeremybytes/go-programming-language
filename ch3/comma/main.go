// Comma inserts commas in a non-negative integer string
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for _, s := range os.Args[1:] {
		_, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			fmt.Printf("%s is not a number", s)
			continue
		}
		fmt.Printf("%s = %s\n", s, comma(s))
	}
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}
