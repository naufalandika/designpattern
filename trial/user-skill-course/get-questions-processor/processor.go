package getquestionsprocessor

import (
	"context"
)

type GetQuestionsProcessor interface {
	Validate(context.Context, *ValidateRequest) error
}

func GetGetQuestionsProcessor(ctx context.Context, req *GetGetQuestionsProcessorRequest) GetQuestionsProcessor {
	switch {
	case req.UserType == PrakerjaUserType && req.CourseType == OnlineCourseType:
		return &PrakerjaOnlineGetQuestionsProcessor{}
	case req.UserType == PrakerjaUserType && req.CourseType == WebinarCourseType:
		return &PrakerjaWebinarGetQuestionsProcessor{}
	default:
		return &BaseProcessor{}
	}
}
