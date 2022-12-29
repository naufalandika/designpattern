package getquestionsprocessor

import (
	"context"
	"fmt"
)

var (
	c = 0
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
	ut := (UserType(c % 2))

	switch ut {
	case NonPrakerjaUserType:
		fmt.Println("usertype: Non Prakerja")
	case PrakerjaUserType:
		fmt.Println("usertype: Prakerja")
	}

	c++

	return ut
}

func getCourseType(req *GetGetQuestionsProcessorRequest) CourseType {
	ct := OnlineCourseType

	switch ct {
	case OnlineCourseType:
		fmt.Println("coursetype: Online")
	}

	return ct
}
