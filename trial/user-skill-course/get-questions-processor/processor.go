package getquestionsprocessor

import (
	"context"
)

type Processor interface {
	Validate(context.Context, *ValidateRequest) error
}

func GetProcessor(ctx context.Context, req *GetProcessorRequest) Processor {
	switch {
	case req.UserType == PrakerjaUserType && req.CourseType == OnlineCourseType:
		return &PrakerjaOnlineProcessor{}
	case req.UserType == PrakerjaUserType && req.CourseType == WebinarCourseType:
		return &PrakerjaWebinarProcessor{}
	default:
		return &BaseProcessor{}
	}
}
