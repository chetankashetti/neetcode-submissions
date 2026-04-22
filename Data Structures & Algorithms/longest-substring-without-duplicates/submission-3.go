func lengthOfLongestSubstring(s string) int {
	length, l := 0, 0
	m := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		char := s[r]
		if lastIdx, ok := m[char]; ok && lastIdx >= l {
			l = lastIdx + 1
		}
		m[char] = r
		if r - l + 1 > length {
			length = r - l + 1
		}
	}
	return length
}
