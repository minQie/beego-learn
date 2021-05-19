package utils

// PS：go 的 sort 包下也有如下各种类型的切片，不过主要目的是为了排序
type StrSlice []string
type IntSlice []int
type Int64Slice []int64

func (s StrSlice) Contains(theStr string) bool {
	for _, str := range s {
		if str == theStr {
			return true
		}
	}
	return false
}

func (s IntSlice) Contains(theI int) bool {
	for _, i := range s {
		if i == theI {
			return true
		}
	}
	return false
}

func (s Int64Slice) Contains(theI int64) bool {
	for _, i := range s {
		if i == theI {
			return true
		}
	}
	return false
}
