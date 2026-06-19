func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    bucket := make([][]int, len(nums)+1)
    for num, count := range freq {
        bucket[count] = append(bucket[count], num)
    }

    result := make([]int, 0, k)
    for i:= len(nums); i >= 0; i-- {
        for _, v := range bucket[i] {
            result = append(result, v)

            if len(result) == k {
                return result
            }
        }
    }

    return result
}