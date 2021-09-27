package slice

type Int64 []int64

func (s Int64) Contain(e int64) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}
	return false
}
