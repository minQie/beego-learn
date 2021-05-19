package form

type Address struct {
	Name string `form:"name"`
	Tag  string `form:"tag"`
}
