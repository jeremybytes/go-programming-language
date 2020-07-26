// Echo3 prints is command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[0:] {
		fmt.Printf("Index %v: %s\n", i, arg)
	}
}

// func main() {
// 	fmt.Println(strings.Join(os.Args[1:], " "))
// }
