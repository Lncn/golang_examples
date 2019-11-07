package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	/* You can convert a time to seconds since epoch using the Unix() func! */
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	/* You can also take a time in seconds since epoch into Time */
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
