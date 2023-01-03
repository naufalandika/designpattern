package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
)

type Service interface {
	GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error)
}

func NewService(accomplishment *accomplishment.Accomplishment) Service {
	return &service{accomplishment}
}

type service struct {
	accomplishment *accomplishment.Accomplishment
}

func (s *service) GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	if err := s.validateGetQuestionsProgress(ctx, req); err != nil {
		return nil, err
	}

	// base get questions logic here

	s.coupledFunction()
	return &GetQuestionsResponse{}, nil
}

func (s *service) validateGetQuestionsProgress(ctx context.Context, req *GetQuestionsRequest) error {
	progress := s.accomplishment.GetUserProgress()

	if progress < 100 {
		return errors.New("user progress < 100")
	}

	return nil
}

func (s *service) coupledFunction() {}
