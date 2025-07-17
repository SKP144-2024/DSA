// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	p := "aaba#aabaab"
	// p := "aabaacaabaad"
	// p := "aabaaabaab"
	lps := []int{0}
	length := 0
	for i := 1; i < len(p); i++ {
		if p[length] == p[i] {
			length += 1
		} else {
			for length != 0 && p[length] != p[i] {
				length = lps[length-1]
			}
			if length != 0 {
				length += 1
			}
		}
		lps = append(lps, length)
	}
	fmt.Println(lps)
}
