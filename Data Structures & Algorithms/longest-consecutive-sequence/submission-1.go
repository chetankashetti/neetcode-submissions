func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}

	cnt := 0
	max := 0
	for i := 0; i < len(nums); i++ {
		j := nums[i]
		for m[j] {
			cnt++
			j++
		}
		if cnt > max {
			max = cnt
		}
		cnt = 0
	}
	return max
}
