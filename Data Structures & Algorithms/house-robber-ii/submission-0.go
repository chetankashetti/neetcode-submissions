func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }
    scenarioA := robLinear(nums[0:len(nums)-1])
    scenarioB := robLinear(nums[1:len(nums)])

    return max(scenarioA, scenarioB)
}

func robLinear(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    if len(nums) == 1 {
        return nums[0]
    }
    if len(nums) == 2 {
        return max(nums[0], nums[1])
    }

    twoHouseBack := nums[0]
    oneHouseBack := max(nums[0], nums[1])
    maxRob := 0
    for i:=2; i<len(nums); i++ {
        maxRob = max(oneHouseBack, twoHouseBack+nums[i])
        twoHouseBack = oneHouseBack
        oneHouseBack = maxRob
    }

    return maxRob
}