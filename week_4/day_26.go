package week4

import (
	"strings"
)

// 394. Decode String https://leetcode.com/problems/decode-string/description/
func DecodeString(s string) string {
	res := []string{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '[' {

		}
		res = append(res, string(s[i]))
	}
	return strings.Join(res, "")
}
