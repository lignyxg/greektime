package source

import "container/list"
func maxSlidingWindow(nums []int, k int) []int {
	dequeue := list.New() // 双端队列中存储数组元素下标
	res := make([]int, 0)
	for i, v := range nums {
		// 窗口向前滑动，如果队列中的最大元素已经不在窗口范围内则弹出
		if i >= k && dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}
		// 考察新进入窗口前端的元素，如果队列不空，且其中的元素比新元素小，则弹出队列中的元素
		for dequeue.Len() > 0 && dequeue.Back().Value.(int) <= v {
			dequeue.Remove(dequeue.Back())
		}
		dequeue.PushBack(i) // 新元素索引入队
		if i >= k-1 {
			j := dequeue.Front().Value.(int)
			res = append(res, nums[j])
		}
	}
	return res
}
