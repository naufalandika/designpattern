package submitposttestprocessor

const (
	PrakerjaUserType    = "prakerja"
	NonPrakerjaUserType = "nonprakerja"
)

const (
	OnlineCourseType  = "online"
	WebinarCourseType = "webinar"
	OfflineCourseType = "offline"
)

type ValidateRequest struct{}

type GetProcessorRequest struct {
	UserType   string
	CourseType string
}
