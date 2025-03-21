package main

import (
	"strings"
)

func isAlNum(b byte) bool {
	if b >= 97 && b <= 122 || b >= 48 && b <= 57 {
		return true
	} else {
		return false
	}

}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	for i, j := 0, len(s)-1; i < j; {
		if isAlNum(s[i]) {
			if isAlNum(s[j]) {
				if s[i] == s[j] {
					i++
					j--
					continue
				} else {
					return false
				}
			} else {
				j--
			}
		} else {
			i++
		}

	}
	return true
}
