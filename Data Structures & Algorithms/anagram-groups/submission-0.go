func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, str := range strs {
		b := []rune(str) 
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		m[key] = append(m[key], str)
	}
	result := [][]string{}
	for _, group := range m {
		result = append(result, group)
	}
	return result
}
