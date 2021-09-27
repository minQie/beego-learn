package slice

type String []string

func (s String) Contain(e string) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}
	return false
}
