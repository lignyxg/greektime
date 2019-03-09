package source

func FirstMissingPositive(nums []int) int {
	for i := 0; i < len(nums);{//处理每个位置
	//每次把一个在数组下表范围内的数字放在对应的下标减1位置
		if nums[i]>0 && nums[i]<len(nums) && nums[nums[i]-1]!=nums[i] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		} else {
			i++
		}
	}

	for i := 0; i < len(nums); i++ {//顺序扫描查找对应位置是不是与数字匹配
		if nums[i] != i+1 {//如果不匹配,则就是这个数字
			return i+1
		}
	}
	//如果所有位置都匹配,返回数组长度+1的数字
	return len(nums)+1
}
