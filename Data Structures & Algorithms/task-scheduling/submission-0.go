func leastInterval(tasks []byte, n int) int {
	freq := make(map[byte]int)
	for _, v := range tasks {
		freq[v]++
	}

	maxFreq := 0
	for _, cnt := range freq {
		if cnt > maxFreq {
			maxFreq = cnt
		}
	}

	countMax := 0 
	for _, cnt := range freq {
		if cnt == maxFreq {
			countMax++
		}
	}

	totalTasks := max(len(tasks), (maxFreq-1)*(n+1)+countMax)
	return totalTasks
}
