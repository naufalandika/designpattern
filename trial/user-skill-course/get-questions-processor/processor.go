package getquestionsprocessor

import (
	"context"

	progressgetter "github.com/naufalandika/designpattern/trial/user-skill-course/progress-getter"
)

var (
	prakerjaOnlineProcessor  *PrakerjaOnlineProcessor
	prakerjaWebinarProcessor *PrakerjaWebinarProcessor
	baseProcessor            *BaseProcessor
)

type Processor interface {
	Validate(context.Context, *ValidateRequest) error
}

func GetProcessor(ctx context.Context, req *GetProcessorRequest) Processor {
	switch {
	case req.UserType == PrakerjaUserType && req.CourseType == OnlineCourseType:
		return prakerjaOnlineProcessor
	case req.UserType == PrakerjaUserType && req.CourseType == WebinarCourseType:
		return prakerjaWebinarProcessor
	default:
		return baseProcessor
	}
}

func init() {
	prakerjaOnlineProcessor = &PrakerjaOnlineProcessor{
		progressGetter: progressgetter.New(),
	}

	prakerjaWebinarProcessor = &PrakerjaWebinarProcessor{}
	baseProcessor = &BaseProcessor{}
}
