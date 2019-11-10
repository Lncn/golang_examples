package main

import (
	"fmt"
	"strconv"
)

func main() {

	/* The second parameter is the base, "0" means infer the base based on input string */
	i, _ := strconv.ParseInt("42", 0, 64)
	fmt.Println(i)
	i, _ = strconv.ParseInt("42", 16, 64)
	fmt.Printf("0x%x\n", i)

	/* ParseInt can parse hex values... */
	i, _ = strconv.ParseInt("0x1ff", 0, 64)
	fmt.Println(i)

	/* I'm not 100% sure where this is useful, but there's a ParseUint.  I guess this is
	   when you want to guarantee the input is non-negative? */
	u, _ := strconv.ParseUint("-1", 0, 8)
	fmt.Println(u)

	/* You can also parse floats... */
	f, _ := strconv.ParseFloat("383.2828773", 64)
	fmt.Println(f)

	/* For convenience to not have to specify a base or bitsize, there's the C equivalent "atoi" */
	x, _ := strconv.Atoi("23836")
	fmt.Println(x)

	/* Obviously you can handle the error too */
	x, e := strconv.Atoi("not an int")
	fmt.Println(e)
}
