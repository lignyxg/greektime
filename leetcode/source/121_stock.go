package source

import "github.com/ethereum/go-ethereum/common/math"

//前i天的最大收益 = max{前i-1天的最大收益, 第i天的价格-前i-1天的最低价格}
func MaxProfit(prices []int) int {
	minPrice, maxProfit := math.MaxInt32,0
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if maxProfit < prices[i]-minPrice {
			maxProfit = prices[i]-minPrice
		}
	}
	return maxProfit
}
