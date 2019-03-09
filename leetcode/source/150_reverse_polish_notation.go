package source

import "strconv"

func EvalRPN(tokens []string) int {
	n := len(tokens)
	if n == 0 {
		return 0
	}
	if n == 1 {
		num, _ := strconv.ParseInt(tokens[0],10,32)
		return int(num)
	}
	ans := 0
	stack := make([]int, n)
	cur := 0 // 模拟栈顶
	for i := 0; i < n; i++ {
		switch tokens[i] {
		case "+":
			x, y := stack[cur-1], stack[cur-2]
			cur -= 2
			ans = x+y
			stack[cur] = ans
			cur++
			break
		case "-":
			x, y := stack[cur-1], stack[cur-2]
			cur -= 2
			ans = x-y
			stack[cur] = ans
			cur++
			break
		case "*":
			x, y := stack[cur-1], stack[cur-2]
			cur -= 2
			ans = x*y
			stack[cur] = ans
			cur++
			break
		case "/":
			x, y := stack[cur-1], stack[cur-2]
			cur -= 2
			ans = x/y
			stack[cur] = ans
			cur++
			break
		default:
			num, _ := strconv.ParseInt(tokens[i],10,32)
			stack[cur] = int(num)
			cur++
		}
	}
	return ans
}
