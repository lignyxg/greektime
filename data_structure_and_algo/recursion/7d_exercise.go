package recursion

func Fibonacci(n int) int {
	if n < 2 {
		if n == 0 {
			return 0
		} else {
			return 1
		}
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Fibonacci2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	a, b, c := 0, 1, 1
	for i := 2; i <= n; i++ {
		c = a+b
		a = b
		b = c
	}
	return c
}

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n*Factorial(n-1)
}
//全排列
var permutations = make([][]int, 0)
func permutation(nums, res []int, n int) {
	if len(res) == n {
		permutations = append(permutations, res)
		return
	}
	for i := 0; i < len(nums); i++ {
		res1 := make([]int, len(res))
		copy(res1, res)
		res1 = append(res1, nums[i])
		nums1 := make([]int, len(nums))
		copy(nums1, nums)
		nums1 = append(nums1[:i], nums1[i+1:]...)
		permutation(nums1, res1, n)
	}
}

func Permutation(nums []int) [][]int {
	res := make([]int, 0)
	permutation(nums, res, len(nums))
	return permutations
}