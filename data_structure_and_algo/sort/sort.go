package sort

import "fmt"

//快速排序
func QuickSort(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	r := len(a) - 1
	quickSort(a, 0, r)
}

func quickSort(a []int, p, r int) {
	if p >= r {
		return
	}

	q := partition2(a, p, r)
	//递归处理枢轴元素分区后的两半
	quickSort(a, p, q-1)
	quickSort(a, q+1, r)
}
//快排分区函数，确定枢轴元素的下标
func partition(a []int, p, r int) int {
	pivot := a[r]
	// 思路：将比pivot小的元素放在pivot前面，大的放在其后面
	i,j := p, r
	for i < j {//前后两个游标i，j
		if a[i] <= pivot {
			i++
		} else {
			a[j] = a[i]
			for a[j] >= pivot && i < j {
				j--
			}
			a[i] = a[j]
		}
	}
	a[i] = pivot
	return i
}

func partition2(a []int, p, r int) int {
	pivot := a[r]
	i, j := p, p
	for j < r {// i,j两个游标分别指向小于pivot的区间和大于pivot的区间
	// i 指向小于pivot区间的下一个元素
		if a[j] < pivot {// 每次有小于pivot的元素, 都将其扔到前面小于pivot的区间
			// swap a[i] with a[j]
			tmp := a[i]
			a[i] = a[j]
			a[j] = tmp
			i++
		}
		j++
	}
	//swap a[i] with a[r]
	a[i] = a[i]^a[r]
	a[r] = a[i]^a[r]
	a[i] = a[i]^a[r]
	return i
}

// 归并排序

func MergeSort(a []int) []int {
	if a == nil || len(a) < 2 {
		return a
	}
	l := len(a)
	return mergeSort(a, 0, l-1)
}

func mergeSort(a []int, p, r int) []int {
	if p == r {
		return a[p:r+1]
	}
	q := (p+r)/2
	part1 := mergeSort(a, p, q)
	part2 := mergeSort(a, q+1, r)
	//a = merge(a[p:q+1], a[q+1:r+1])//给形参赋值是没用的
	return merge(part1, part2)
}
//合并两个有序数组
func merge(a, b []int) []int {
	res := make([]int, len(a)+len(b))
	i, j, la, lb := 0, 0, len(a), len(b)
	k := 0
	for i < la && j < lb {
		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
				res[k] = b[j]
				j++
		}
		k++
	}
	for i < la {
		res[k] = a[i]
		i++
		k++
	}
	for j < lb {
		res[k] = b[j]
		j++
		k++
	}
	return res
}

func MergeSort2(a []int) {
	if a == nil || len(a) < 2 {
		return
	}
	l := len(a)
	mergeSort2(a, 0, l-1)
}

func mergeSort2(a []int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r)/2
	mergeSort2(a, p, q)
	mergeSort2(a, q+1, r)
	merge2(a[p:r+1], a[p:q+1], a[q+1:r+1])
	fmt.Println(a[p:r+1])
}
//将a，b合并到sum
func merge2(sum, a, b []int) {
	p1 := make([]int, len(a))
	p2 := make([]int, len(b))
	copy(p1, a)
	copy(p2, b)

	i, j, k := 0,0,0
	for i < len(a) && j < len(b) {
		if p1[i] < p2[j] {
			sum[k] = p1[i]
			i++
			k++
		} else {
				sum[k] = p2[j]
				j++
				k++
		}
	}
	for i < len(p1) {
		sum[k] = p1[i]
		i++
		k++
	}
	for j < len(p2) {
		sum[k] = p2[j]
		j++
		k++
	}
}

//计数排序(只能对非负整数排序，否则需要处理成非负整数)
//适用于元素范围不大的情况
//记录数组中每个元素出现的次数
func CountSort(a []int) {
	maxElem := 0
	for i := 0; i < len(a); i++ {
		if maxElem < a[i] {
			maxElem = a[i]
		}
	}// 元素的数据范围是0～maxElem
	//计数数组
	elems := make([]int, maxElem+1)
	for i := 0; i < len(a); i++ {
		elems[a[i]]++
	}
	//元素计数累加
	for i := 1; i < len(elems); i++ {
		elems[i] += elems[i-1]
	}

	//根据计数把元素填回原数组，使用原数组的副本，以免赋值干扰
	//elems[i]中记录的是: 小于等于i的元素数，也即包括i在内，按升序目前共有多少个元素
	//因此元素i在排序后数组中的位置就是第i个元素，下标为`i-1`
	b := make([]int, len(a))
	copy(b, a)
	for j := len(a)-1; j >= 0; j-- {
		index := elems[b[j]] - 1
		a[index] = b[j]
	}
}