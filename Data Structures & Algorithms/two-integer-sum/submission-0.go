func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        if idx, ok := numMap[target-num]; ok {
            return []int{idx, i}
        }
        numMap[num] = i
    }
    return []int{}
}
