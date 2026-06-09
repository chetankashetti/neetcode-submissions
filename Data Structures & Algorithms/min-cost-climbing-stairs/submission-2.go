func minCostClimbingStairs(cost []int) int {
    if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}

	oneStepBack := 0
	twoStepBack := 0
	minCost := 0
	for i := 2; i <= len(cost); i++ {
		minCost = min(oneStepBack+cost[i-1], twoStepBack+cost[i-2])
		twoStepBack = oneStepBack
		oneStepBack = minCost
	}

	return minCost
}
