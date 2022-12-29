package getquestionsprocessor

import (
	"context"
	"time"
)

type GetQuestionsProcessor interface {
	Validate(context.Context, *ValidateRequest) error
}

func GetGetQuestionsProcessor(ctx context.Context, req *GetGetQuestionsProcessorRequest) GetQuestionsProcessor {
	ut := getUserType(ctx)
	ct := getCourseType(req)

	switch {
	case ut == PrakerjaUserType && ct == OnlineCourseType:
		return &PrakerjaOnlineGetQuestionsProcessor{}
	}

	return nil
}

func getUserType(ctx context.Context) UserType {
	return (UserType(time.Now().Nanosecond()) % 2)
}

func getCourseType(req *GetGetQuestionsProcessorRequest) CourseType {
	return OnlineCourseType
}
