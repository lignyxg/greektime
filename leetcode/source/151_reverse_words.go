package source

import (
	"strings"
)

func ReverseWords(s string) string {
	b := []byte(strings.TrimSpace(s))
	if len(b) <= 1 {
		return string(b)
	}
	end := 0
	count := 0
	//预处理去掉单词之间的多余空白
	for i := 0; i < len(b); i++ {
		if b[i] != byte(' ') {
			b[end] = b[i]
			end++
			count = 0
		} else {
			if count == 0 {
				count++
				b[end] = ' '
				end++
			} else {
					continue
			}
		}
	}
	b = b[:end]
	reverse(b)
	j := 0
	for i := 0; i <= end; i++ {
		if i == end {
			reverse(b[j:i])
			break
		}
		if b[i] == byte(' ') {
			reverse(b[j:i])
			j = i+1
		}
	}
	return string(b)
}

func reverse(b []byte) {
	for l, h := 0,len(b)-1; l < h; {
		tmp := b[l]
		b[l] = b[h]
		b[h] = tmp
		l++
		h--
	}
}
