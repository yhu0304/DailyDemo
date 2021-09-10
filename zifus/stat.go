package zifus

import (
	"strconv"
	"strings"
)

func Stat64String(number int64, suffix string) (s string) {
	if number == 0 {
		s = "-" + suffix
		return
	}
	if number < 10000 {
		s = strconv.FormatInt(number, 10) + suffix
		return
	}
	if number < 100000000 {
		s = strconv.FormatFloat(float64(number)/10000, 'f', 1, 64)
		return strings.TrimSuffix(s, ".0") + "万" + suffix
	}
	s = strconv.FormatFloat(float64(number)/100000000, 'f', 1, 64)
	return strings.TrimSuffix(s, ".0") + "亿" + suffix
}
