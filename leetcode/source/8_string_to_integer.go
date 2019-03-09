package source

import (
	"github.com/ethereum/go-ethereum/common/math"
	"strings"
)

func MyAtoi(str string) int {
	flag := true // 正数还是负数?
	b := []byte(strings.TrimLeft(str, " ")) // 去掉开头的空白
	if len(b) == 0 {
		return 0
	}
	if b[0] == byte('+') {
		flag = true
		b = b[1:]
	} else if b[0] == byte('-') {
		flag = false
		b = b[1:]
	} else if b[0]<48 || b[0]>57 { // 开头为非法字符
		return 0
	}
	nb := make([]uint8, len(b))
	end := 0
	mark := true //标志是否找到了第一个非零数字
	for i := 0; i < len(b); i++ {
		if b[i] == 48 && mark {//跳过开头的0
			continue
		}
		if b[i] >= 48 && b[i] <= 57 {
			mark = false
			nb[end] = b[i] - 48
			end++
		} else {
			break
		}
	}
	res := 0
	nb = nb[:end] // 截取有效数字
	if end > 10 { // 一定超过了int32的最大表示范围
		if flag {
			return math.MaxInt32
		} else {
			return math.MinInt32
		}
	}
	maxValue := math.MaxInt32
	k := 1
	for i := end-1; i >= 0; i-- { //累加数字
		res = res + int(nb[i])*k
		maxValue -= int(nb[i])*k
		if maxValue < 0 { // 超过了int32的最大表示范围
			break
		}
		k *= 10
	}
	if !flag {//负数
		if maxValue < -1 {
			return math.MinInt32
		}
		return -res
	}

	if maxValue < 0 {
		return math.MaxInt32
	}
	return res
}
