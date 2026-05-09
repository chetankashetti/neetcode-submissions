func characterReplacement(s string, k int) int {
    maxFreq, maxLen, l := 0, 0, 0
    m := make(map[byte]int)
    for r := 0; r < len(s); r++ {
        m[s[r]]++

        maxFreq = max(maxFreq, m[s[r]])

        window := r - l + 1

        if window -maxFreq > k {
            m[s[l]]--
            l++
        } 

        maxLen = max(maxLen, r-l+1)
    }
    return maxLen
}
