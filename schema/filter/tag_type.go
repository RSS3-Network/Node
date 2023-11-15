package filter

type Type interface {
	Name() string
	Tag() Tag
}
