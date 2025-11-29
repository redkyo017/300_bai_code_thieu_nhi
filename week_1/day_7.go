package week1

// 76. Minimum Window Substring https://leetcode.com/problems/minimum-window-substring/description/
func MinWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	freq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		freq[t[i]]++
	}

	matched := 0
	required := len(freq)

	minLen := len(s) + 1
	startIndex := 0

	left := 0

	windowCount := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		charR := s[right]

		windowCount[charR]++

		if countT, exists := freq[charR]; exists {
			if windowCount[charR] == countT {
				matched++
			}
		}

		for matched == required {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
				startIndex = left
			}

			charL := s[left]
			windowCount[charL]--

			if countT, exists := freq[charL]; exists {
				if windowCount[charL] < countT {
					matched--
				}
			}

			left++
		}
	}

	if minLen > len(s) {
		return ""
	}

	return s[startIndex : startIndex+minLen]
}
