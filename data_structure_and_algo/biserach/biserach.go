package biserach

import "sync"

//在循环有序数组中查找特定元素
//循环有序数组，如 [4,5,6,1,2,3]
//假设数组有序部分按照升序排列，则前半部分的所有元素一定大于后半部分的所有元素
//思路：每次二分查找，会把数组分成两半，一半是有序数组，一半依然是循环数组，
//这可以通过mid与首元素对比得知，然后确定target是否会在有序数组中还是在循环数组中，
//对应舍弃target不会出现的那一半数组，然后继续查找
func FindInCyclingArray(nums []int, length, target int) int {
	low, high := 0, length-1
	for low <= high {
		mid := low +((high-low)>>1)
		if nums[mid] < target {
			if nums[mid] >= nums[low] { //数组前半部分有序，后半部分为循环数组
				//目标元素在后半个区间, 因为target已经比有序区间最大的元素还要大
				low = mid + 1
			} else {//数组后半部分有序
				if nums[high] >= target { // 判断元素是否在有序区间
					low = mid + 1
				} else {
					high = mid - 1
				}
			}
		} else if nums[mid] > target {
			if nums[mid] >= nums[low] { //前半部分有序
				if nums[low] <= target { // 元素是否在有序区间中?
					high = mid - 1
				} else {
					low = mid + 1
				}
			} else {//后半部分有序,后半部分所有值都比nums[mid]更大
				high = mid - 1
			}
		} else {
			return mid
		}
	}
	return -1
}

//变体1：找到第一个等于某个值的元素
func FindFirstEqual(nums []int, target int) int {
	var m sync.Map
}
