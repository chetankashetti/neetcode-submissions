func findKthLargest(nums []int, k int) int {
	targetIdx := len(nums) - k

	var quickSelect func(l int, r int) int
	quickSelect = func(l int, r int) int {
		randomIdx := l + rand.Intn(r-l+1)
		nums[randomIdx], nums[r] = nums[r], nums[randomIdx]
		pivot := nums[r]
		p := l


		for i := l; i < r; i++ {
			if nums[i] <= pivot {
				nums[i], nums[p] = nums[p], nums[i]
				p++
			}
		}
		nums[p], nums[r] = nums[r], nums[p]

		if p > targetIdx {
			return quickSelect(l, p-1)
		} else if p < targetIdx {
			return quickSelect(p+1, r)
		} else {
			return nums[p]
		}
	}

	return quickSelect(0, len(nums)-1)
}
