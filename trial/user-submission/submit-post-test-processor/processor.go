package submitposttestprocessor

import "context"

type Processor interface {
	Validate(context.Context, *ValidateRequest) error
	PostSubmission(context.Context, *PostSubmissionRequest) error
}

func GetProcessor(ctx context.Context, req *GetProcessorRequest) Processor {
	switch {
	case req.UserType == NonPrakerjaUserType && req.CourseType == OnlineCourseType:
		return &NonPrakerjaOnlineProcessor{}
	case req.UserType == PrakerjaUserType && req.CourseType == OnlineCourseType:
		return &PrakerjaOnlineProcessor{}
	default:
		return &BaseProcessor{}
	}
}
