// Sha256 compares 2 crypto values
package main

import (
	"crypto/sha256"
	"fmt"
)

func getBit(b byte, bitPos int) byte {
	r := b << (9 - bitPos)
	r = r >> 7
	return r
}

func bitDiff(x, y []byte) int {
	var count int
	for i := 0; i < len(x) && i < len(y); i++ {
		for b := 1; b <= 8; b++ {
			if getBit(x[i], b) == getBit(y[i], b) {
				count++
			}
		}
	}
	return count
}

func main() {

	//testing the bit comparisons
	// b1 := []byte{1}
	// b2 := []byte{7}

	// fmt.Printf("%08b\n", b1)
	// fmt.Printf("%08b\n", b2)
	// fmt.Printf("%d\n", bitDiff(b1, b2))

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%08b\n%08b\n%t\n%T\n", c1, c2, c1, c2, c1 == c2, c1)

	fmt.Printf("%d matching bits out of %d\n", bitDiff(c1[:], c2[:]), len(c1)*8)

}
