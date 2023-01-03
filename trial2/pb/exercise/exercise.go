package exercise

type Exercise struct{}

func New() *Exercise {
	return &Exercise{}
}

func (e *Exercise) Submit() int32 {
	return 80
}
