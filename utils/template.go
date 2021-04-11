package utils

import "strings"

func GetBatchInsertConditionStr(fieldSum int) string {
	if fieldSum == 0 {
		panic(`zero sum`)
	}

	var b strings.Builder
	b.Grow((fieldSum*2 + 1) * 8)

	b.WriteString(`(?`)
	for index2 := 1; index2 < fieldSum; index2++ {
		b.WriteString(`,?`)
	}
	b.WriteString(`)`)

	return b.String()
}
