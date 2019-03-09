package source

//状态转移方程: F(amount) = min{F(amount-y)+1},y是其中一种零钱
func CoinChange(coins []int, amount int) int {
	//状态数组
	s := make([]int, amount+1)
	s[0] = 0
	for i := 1; i <= amount; i++ {
		min := amount+1
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 && s[i-coins[j]] != -1 {
				if min > s[i-coins[j]]+1 {
					min = s[i-coins[j]]+1
				}
			}
		}
		if min == amount+1 {
			s[i] = -1
		} else {
			s[i] = min
		}
	}
	return s[amount]
}
