// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	// p := "aaba#aabaab"
	// Ans : [0 1 0 1 0 1 2 3 4 2 3]
	p := "aabaacaabaad"
	// Ans : [0 1 0 1 2 0 1 2 3 4 5 0]
	// p := "aabaaabaab"
	// Ans : [0 1 0 1 2 2 3 4 5 3]
	lps := make([]int, len(p))
	l := 0
	for i := 1; i < len(p); i++ {
		for l > 0 && p[l] != p[i] {
			l = lps[l-1]
		}
		if p[l] == p[i] {
			l += 1
			lps[i] = l
		}
	}
	fmt.Println(lps)
}
