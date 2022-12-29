package getquestionsprocessor

type UserType int

const (
	NonPrakerjaUserType UserType = iota
	PrakerjaUserType
)

type CourseType int

const (
	OnlineCourseType CourseType = iota
)

type ValidateRequest struct{}

type GetGetQuestionsProcessorRequest struct{}
