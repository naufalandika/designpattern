package userskillcourse

import (
	"context"

	getquestionsprocessor "github.com/naufalandika/designpattern/trial2/user-skill-course/get-questions-processor"
)

type Service interface {
	GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	processor := getquestionsprocessor.GetProcessor(ctx, &getquestionsprocessor.GetProcessorRequest{
		UserType:   req.UserType,
		CourseType: req.CourseType,
	})

	if err := processor.Validate(ctx, &getquestionsprocessor.ValidateRequest{}); err != nil {
		return nil, err
	}

	// base get questions logic here

	s.coupledFunction()
	return &GetQuestionsResponse{}, nil
}

func (s *service) coupledFunction() {}
