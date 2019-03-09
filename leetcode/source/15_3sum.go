package source

import (
	"sort"
)

var res = make([][]int, 0)
func ThreeSum(nums []int) [][]int {
	//res := make([][]int, 0)
	items := make([]int, 0)
	three(nums, items)
	return res
}
//无法去重
func three(nums []int, items []int) {
	//fmt.Println("nums:", nums)
	if len(items) == 3 {
		sum := 0
		for _, x := range items {
			sum += x
		}
		if sum == 0 {
			//fmt.Println(items)
			res = append(res, items)
		}
		return
	}

	for i := 0; i < len(nums); i++ {
		c := make([]int, len(items))
		copy(c, items)
		c = append(c, nums[i])
		three(nums[i+1:], c)
	}
}

func ThreeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] { // 跳过重复解
			sum, l, r := -nums[i], i+1, len(nums)-1
			for l < r {
				if nums[l]+nums[r] == sum {
					result = append(result, []int{nums[i], nums[l], nums[r]})
					for l < r && nums[l] == nums[l+1] {// 跳过重复解
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					l++
					r--
				} else if nums[l]+nums[r] < sum {
					for l < r && nums[l] == nums[l+1] {// 跳过重复解
						l++
					}
					l++
				} else {
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				}
			}
		}
	}

	return result
}