package main

import (
	"fmt"
	"sort"
)

// Pretend we want to sort a string slice by length instead of by default
// alphabetically. To do this, we need to implement the `sort` library's
// interface (`sort.Interface`), which includes the three methods:
//	`Len`, `Less`, `Swap`
type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "kiwi", "banana"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
