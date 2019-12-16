package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	/* This function creates a new buffered reader on stdin and converts
	 * everything ToUpper before outputting to stdout (until newline) */
	scanner := bufio.NewScanner(os.Stdin)

	/* Scan runs until EOF */
	for scanner.Scan() {
		upper := strings.ToUpper(scanner.Text())
		fmt.Println(upper)
	}

	/* I think the only way to hit an error here is to input file without EOF? */
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
