package source

import "sort"

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closestNum := nums[0]+nums[1]+nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			l, r := i+1, len(nums)-1
			for l < r {
				threeSum := nums[i]+nums[l]+nums[r]
				if abs(threeSum-target) < abs(closestNum-target) {
					closestNum = threeSum

				}
				if threeSum > target {
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				} else if threeSum < target {
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
				} else { // 与target相等肯定是最接近的
					return target
				}
			}
		}
	}
	return closestNum
}

func abs(num int) int {
	if num >= 0 {
		return num
	} else {
		return -num
	}
}