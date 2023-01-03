package accomplishment

type Accomplishment struct{}

func New() *Accomplishment {
	return &Accomplishment{}
}

func (a *Accomplishment) GetUserProgress() int32 {
	return 80
}
