package utils

import (
	"github.com/pkg/errors"
	"regexp"
	"strings"
	"time"
)

func GetDateString() string {
	return DateTimeToDateString(time.Now())
}

// Time(2020-01-01 00:00:00) → '2020-01-01'
func DateTimeToDateString(dateTime time.Time) string {
	return dateTime.Format("2006-01-02")
}

// string(2020-01-01 00:00:00) → 2020-01-01
// 如果参数不是指定的格式，将会返回 error
func DateTimeStringToDateString(dateTimeString string) (string, error) {
	// 校验参数格式
	isMatch := regexp.MustCompile("\\d{4}-\\d{2}-\\d{2} \\d{2}:\\d{2}:\\d{2}").Match([]byte(dateTimeString))
	if !isMatch {
		return "", errors.Errorf("错误的时间格式：%s", dateTimeString)
	}

	return strings.Split(dateTimeString, " ")[1], nil
}
