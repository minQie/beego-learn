package utils

import (
	"bytes"
)

// 获取 sql 需要的参数占位符(Parameter Markers)
func OrmJoinRepeat(n int) string {
	return JoinRepeat("?", ",", n)
}

// join `str` with `sep` for `l` times
func JoinRepeat(str, sep string, l int) string {
	if l < 1 {
		return ""
	}
	if l == 1 {
		return str
	}

	var (
		n      = l*(len(str)+len(sep)) - len(sep)
		sepStr = sep + str
	)
	buf := bytes.Buffer{}
	buf.Grow(n)
	buf.WriteString(str)
	for i := 1; i < l; i++ {
		buf.WriteString(sepStr)
	}
	return buf.String()
}
