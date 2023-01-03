package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
	"github.com/naufalandika/designpattern/trial2/pb/exercise"
)

type Service interface {
	GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error)
}

func NewService(accomplishment *accomplishment.Accomplishment, exercise *exercise.Exercise) Service {
	return &service{
		accomplishment: accomplishment,
		exercise:       exercise,
	}
}

type service struct {
	accomplishment *accomplishment.Accomplishment
	exercise       *exercise.Exercise
}

func (s *service) GetQuestions(ctx context.Context, req *GetQuestionsRequest) (*GetQuestionsResponse, error) {
	if err := s.validateGetQuestionsProgress(ctx, req); err != nil {
		return nil, err
	}

	s.exercise.GetQuestions()

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
