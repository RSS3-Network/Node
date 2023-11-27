package filter

type Chain interface {
	Network() Network
	String() string
	ID() uint64
	FullName() string
}
