package userskillcourse

import (
	"context"

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
	processor := getProcessor(&getProcessorRequest{})

	if err := processor.validateGetQuestionsProgress(ctx, &validateGetQuestionsProgressRequest{}); err != nil {
		return nil, err
	}

	s.exercise.GetQuestions()

	s.coupledFunction()
	return &GetQuestionsResponse{}, nil
}

func (s *service) coupledFunction() {}
