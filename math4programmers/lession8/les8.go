package main

import "fmt"

//输出从n个数中选出m个数的所有组合
func findCombine(nums, result []int, m int) {
	if len(result) == m {
		fmt.Println(result)
		return
	}
	for i := 0; i < len(nums); i++ {
		nextLevelRes := make([]int, len(result))
		copy(nextLevelRes, result)
		nextLevelRes = append(nextLevelRes, nums[i])
		findCombine(nums[i+1:],nextLevelRes,m)
	}
}

//从n个人中分别抽出m1个一等奖，m2个二等奖，m3个三等奖的所有组合
func lottery() {

}


func main() {
	n := []int{1,2,3,4,5}
	res := make([]int, 0)
	findCombine(n, res, 4)
}