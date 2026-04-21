func topKFrequent(nums []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	buckets := make([][]int, len(nums)+1)
	for num, cnt := range m {
		buckets[cnt] = append(buckets[cnt], num)
	}

	res := []int{}
	for f:=len(buckets)-1; f >= 0 && len(res) < k; f-- {
		if len(buckets[f]) == 0 {
			continue
		}
		sort.Ints(buckets[f])
		for _, num := range buckets[f] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}
	return res
}
