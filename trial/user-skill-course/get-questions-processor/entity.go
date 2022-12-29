package getquestionsprocessor

type UserType int

const (
	NonPrakerjaUserType = iota
	PrakerjaUserType
)

type CourseType int

const (
	OnlineCourseType = iota
)

type ValidateRequest struct{}

type GetGetQuestionsProcessorRequest struct{}
