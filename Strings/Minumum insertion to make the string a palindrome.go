package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalidromicSubsequence(s string) int {
	n := len(s)
	prevRow := make([]int, n+1)
	sFwd := []rune(s)
	sRev := []rune(s)
	for i := 0; i < n/2; i++ {
		sRev[i], sRev[n-i-1] = sRev[n-i-1], sRev[i]
	}

	for i := 1; i <= n; i++ {
		currRow := make([]int, n+1)
		for j := 1; j <= n; j++ {
			if sFwd[i-1] == sRev[j-1] {
				currRow[j] = 1 + prevRow[j-1]
			} else {
				currRow[j] = max(currRow[j-1], prevRow[i])
			}
		}
		prevRow = currRow
		// fmt.Println(prevRow)
	}

	return prevRow[n]
}

func minInsertionsToPalindrome(s string) int {
	return len(s) - longestPalidromicSubsequence(s)
}

// ✅ Test runner
func runTestCase(s string, expected int) {
	output := minInsertionsToPalindrome(s)
	fmt.Printf("Input: \"%s\"\n", s)
	fmt.Printf("Expected: %d\n", expected)
	fmt.Printf("Your Output: %d\n", output)

	if output == expected {
		fmt.Println("✅ PASS")
	} else {
		fmt.Println("❌ FAIL")
	}
	fmt.Println("----------------------------------")
}

func main() {
	// 🧪 Testcases

	runTestCase("a", 0)       // already a palindrome
	runTestCase("ab", 1)      // make it "aba" or "bab"
	runTestCase("abc", 2)     // make it "cbabc" or similar
	runTestCase("race", 3)    // "ecarace"
	runTestCase("aebcbda", 2) // make it "adabcbda"
	runTestCase("abcdba", 1)  // insert 1 to fix
	runTestCase("aaaa", 0)    // already palindrome
	runTestCase("aabb", 2)    // minimal insertions
	runTestCase("abca", 1)    // "abcba"
	runTestCase("", 0)        // empty string is palindrome
	runTestCase("abcd", 3)    // "dcbabcd"

	// 🔄 Add more custom stress/edge cases if needed
}
