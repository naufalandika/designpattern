package usersubmission

import (
	"context"
	"fmt"

	submitposttestprocessor "github.com/naufalandika/designpattern/trial/user-submission/submit-post-test-processor"
)

type Service interface {
	SubmitPostTest(context.Context, *SubmitPostTestRequest) (*SubmitPostTestResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) SubmitPostTest(ctx context.Context, req *SubmitPostTestRequest) (*SubmitPostTestResponse, error) {
	processor := submitposttestprocessor.GetProcessor(ctx, &submitposttestprocessor.GetProcessorRequest{
		UserType:   req.UserType,
		CourseType: req.CourseType,
	})

	processor.Validate(ctx, &submitposttestprocessor.ValidateRequest{})

	fmt.Println("base logic submit post test")

	go processor.PostSubmission(ctx, &submitposttestprocessor.PostSubmissionRequest{})

	return nil, nil
}
