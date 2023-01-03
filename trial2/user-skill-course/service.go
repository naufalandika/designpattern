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
	if err := s.getQuestionsValidation(ctx, req); err != nil {
		return nil, err
	}

	// base get questions logic here

	s.coupledFunction()
	return &GetQuestionsResponse{}, nil
}

func (s *service) getQuestionsValidation(ctx context.Context, req *GetQuestionsRequest) error {
	return nil
}

func (s *service) coupledFunction() {}
