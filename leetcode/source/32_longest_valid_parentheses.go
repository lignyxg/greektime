package source

//方法1
//正序倒序各扫描一次，取两次得到的最大值
//此方法利用了字符串中只包含'('与')'一种括号的特点，因此可以简单使用+1和-1来计数
//也就是匹配的结果与'('和')各自的总数量有关（记为a和b）:
// a = b,则初始串一定是有效串
// a < b 或 a > b,则是初始串的某个子串有效
//每次扫描记录最大的有效长度
//从左至右扫描：一旦遇到某个字串中a>b的情况就不能有效记录有效串的最大长度，如"(()"
//但a>b也就是b<a，再从右至左扫描一遍就可以记录有效长度了
//e.g.: "()(()()"
//从左至右：得到长度为2
//从右至左：得到长度为4

func LongestValidParentheses(s string) int {
	l1 := scan(s, 0, len(s), 1, '(')
	l2 := scan(s, len(s)-1, -1, -1, ')')
	if l1 > l2 {
		return l1
	} else {
		return l2
	}
}

func scan(s string, i, end, flag int, c byte) int {
	max, validLen, currentLen, sum:= 0 ,0, 0, 0
	for ; i != end; i += flag {
		if s[i] == c {
			sum++
		} else {
			sum--
		}
		currentLen++
		if sum < 0 {//当前已扫描的子串为无效串
			if max < validLen {
				max = validLen //记录当前最大的有效长度
			}
			validLen, currentLen, sum = 0 ,0, 0 //重新开始计数
		} else if sum == 0 {//找到了一个有效的匹配串
			validLen = currentLen
		}
	}
	if max < validLen {
		max = validLen
	}
	return max
}

//方法2
//改造括号匹配的算法，栈数据类型使用结构体{index, character}
//遍历完成字符串之后，查看栈内容，求出栈中相邻元素的最大差值即可
//栈内容：[{3,')'},{7,'('}]，假如字符串长度为18
//可得匹配子串有三段，0～2，4～6，8～15，则最大匹配长度为8
func LongestValidParentheses2(s string) int {
	//TODO
	return 0
}

//方法3
//使用一个栈和一个辅助数组，栈用于记录匹配括号的索引下标，
// 数组用于记录遍历到当前位置时，若发生匹配，符合条件的最长括号串长度
//
func LongestValidParentheses3(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	stack := make([]int, n) // 模拟栈
	cur := 0 // 栈顶指针
	mark := make([]int, n) // 标记数组
	max := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' { // 左括号索引入栈
			stack[cur] = i
			cur++
			mark[i] = 0 // 左括号记录当前子串长度为0
		} else { // 右括号索引不必入栈
			if cur > 0 { // 当前栈内有左括号
				cur-- // 模拟退栈
				mark[i] = i-stack[cur]+1 // 右和左括号索引相减得到当前长度
				// stack[cur]是与当前右括号匹配的左括号下标
				// 此时需要查看此左括号紧邻前面，即`stack[cur]-1`位置，是否还有连续的符合条件的子串
				// 如果有则合并，可以借助mark数组来判断
				if stack[cur]-1 > 0 && mark[stack[cur]-1] != 0 {//找到紧邻的符合条件的子串
					mark[i] = mark[i] + mark[stack[cur]-1]
				}
			} else {
				//栈为空，即此时没有匹配的左括号
				mark[i] = 0
			}
			if max < mark[i] {
				max = mark[i]
			}
		}
	}
	return max
}

//方法4
//使用栈和辅助数组，栈用于记录括号匹配索引，数组初始化为全0
//每当遇到匹配的括号，就将数组中匹配的左右括号下标之间的值全置为1
//括号匹配完成后，扫描数组，累加连续的1的最大数量即可
func LongestValidParentheses4(s string) int {
	return 0
}

type Stack struct {
	N int
	inner []interface{}
}

func NewStack(n int) *Stack {
	if n < 0 {
		return nil
	}
	return &Stack{
		0,
		make([]interface{}, n),
	}
}
func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return -1
	}
	s.N--
	return s.inner[s.N]
}
func (s *Stack) Push(n interface{}) {
	if s.IsFull() {
		return
	}
	s.inner[s.N] = n
	s.N++
}
func (s *Stack) IsEmpty() bool {
	return s.N == 0
}
func (s *Stack) IsFull() bool {
	return s.N == len(s.inner)
}
func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return -1
	}
	return s.inner[s.N-1]
}