func longestPalindrome(s string) string {
	maxLength, start, end := 0, 0, 0
    for i:=0; i<len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		length := max(len1, len2)
		if length > maxLength {
			maxLength = length
			start = i - (maxLength - 1) / 2
			end = i + (maxLength) / 2
		}
	}
	return s[start:end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return right - left - 1
}
