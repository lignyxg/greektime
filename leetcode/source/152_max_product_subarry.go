package source

import "math"

func MaxProduct(nums []int) int {
	max := nums[0]
	cum := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			cum = 1
			if max < 0 {
				max = 0
			}
			continue
		}
		cum *= nums[i]
		bigger := math.Max(float64(cum), float64(nums[i]))
		if max < int(bigger) {
			max = int(bigger)
		}
	}
	max2 := nums[len(nums)-1]
	cum = 1
	for j := len(nums)-1; j >= 0; j-- {
		if nums[j] == 0 {
			cum = 1
			if max2 < 0 {
				max2 = 0
			}
			continue
		}
		cum *= nums[j]
		bigger := math.Max(float64(cum), float64(nums[j]))
		if max2 < int(bigger) {
			max2 = int(bigger)
		}
	}
	if max2 > max {
		max = max2
	}
	return max
}

func MaxProduct2(nums []int) int {
	max, imax, imin := math.MinInt32,1,1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = int(math.Max(float64(imax*nums[i]),float64(nums[i])))
		imin = int(math.Min(float64(imin*nums[i]), float64(nums[i])))
		max = int(math.Max(float64(max), float64(imax)))
	}
	return max
}
