func coinChange(coins []int, amount int) int {
    infeasible := amount + 1 
    dp := make([]int, amount+1)
    for i:=1; i<=amount;i++ {
        dp[i] = infeasible
    }

    dp[0] = 0

    for i:=1; i<=amount; i++ {
        for _,c := range coins {
            if i >= c {
                if dp[i-c]+1 < dp[i] {
                    dp[i] = dp[i-c] + 1
                }
            }
        }
    }
    if dp[amount] == infeasible {
        return -1
    }
    return dp[amount]
}