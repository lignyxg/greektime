package source

import "fmt"

//递归方式求解
//leetcode运行结果与本地运行结果不一致
var matched = false

func IsMatch(s string, p string) bool {
	if len(s) == 0 && p == ".*" {
		return true
	}
	match([]byte(s), []byte(p), 0, 0)
	return matched
}

func match(s, p []byte, si, pj int) {
	if matched {
		return
	}
	if pj == len(p) {
		if si == len(s) {
			matched = true
		}
		return
	}
	if p[pj] == byte('.') { //匹配单个字符
		match(s, p, si+1, pj+1)
	} else if p[pj] == byte('*') {
		match(s,p,si,pj+1) //匹配0个前面的字符
		fmt.Println("si:",si,"pj:",pj)
		//匹配1个或多个前面的字符
		for i := 0;i < len(s)-si && (p[pj-1] == s[si+i] || p[pj-1] == byte('.'));i++{
			match(s,p,si+i+1,pj+1)
		}
	} else if si<len(s) && p[pj]==s[si]{
		match(s, p, si+1, pj+1)
	}
}

//动态规划方式求解
func IsMatch2(s string, p string) bool {
	ls, lp := len(s), len(p)
	state := make([][]bool, ls+1)
	for k := range state {
		state[k] = make([]bool, lp+1)
	}
	state[0][0] = true
	for i := 0; i <= ls; i++ { // i==0时，代表s为空
		for j := 1; j <= lp; j++ { // p不能为空
			if p[j-1] == '*' {
				//即*之前的字符重复0次后匹配（等价于dp[i][j-2]）或者重复一次后匹配（等价于dp[i-1][j])
				state[i][j] = state[i][j-2] || (i>0 && state[i-1][j] && (s[i-1]==p[j-2] || p[j-2]=='.'))
			} else {
				//非*结尾时，对比最后一位是否可匹配，如果匹配则当前dp[i][j]等价于dp[i-1][j-1]
				state[i][j] = i>0 && state[i-1][j-1] && (s[i-1]==p[j-1] || p[j-1]=='.')
			}
		}
	}
	return state[ls][lp]
}