func maxArea(heights []int) int {
	n := len(heights)
	l, r := 0, n - 1
	maxArea := 0
	for l < r {
		area := 0
		if heights[l] < heights[r] {
			area = heights[l] * (r-l)
			l++
		} else {
			area = heights[r] * (r-l)
			r--
		}
		if area > maxArea {
			maxArea = area
		} 
	}
	return maxArea
}
