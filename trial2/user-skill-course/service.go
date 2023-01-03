package userskillcourse

import (
	"context"
)

type Service interface {
	GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	processor := newGetQuestionsProcessor()

	if err := processor.validate(ctx, &getQuestionsProcessorValidateRequest{}); err != nil {
		return nil, err
	}

	// base get questions logic here

	s.coupledFunction()
	return &GetQuestionsResponse{}, nil
}

func (s *service) coupledFunction() {}
