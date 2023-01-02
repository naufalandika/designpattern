package getquestionsprocessor

const (
	NonPrakerjaUserType = "nonprakerja"
	PrakerjaUserType    = "prakerja"
)

const (
	OnlineCourseType  = "online"
	WebinarCourseType = "webinar"
)

type ValidateRequest struct{}

type GetGetQuestionsProcessorRequest struct {
	UserType   string
	CourseType string
}
