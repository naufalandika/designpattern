package exercise

type Exercise struct{}

func New() *Exercise {
	return &Exercise{}
}

func (e *Exercise) Submit() int32 {
	return 80
}

type Question struct{}

type Questions []Question

func (e *Exercise) GetQuestions() Questions {
	return []Question{}
}
