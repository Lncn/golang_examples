package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "Imma let you hash this string"

	hash := sha1.New()

	/* The Write function expects bytes, so coerce everything to a byte slice */
	hash.Write([]byte(str))

	/* Sum returns the finalized hash as a byte slice.  You can pass the last byte (slice)
	   into the Sum as its parameter, but not usually necessary */
	fmt.Printf("SHA1 of '%s': 0x%x\n", str, hash.Sum(nil))
}
