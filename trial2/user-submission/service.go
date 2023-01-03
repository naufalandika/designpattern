package usersubmission

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
	"github.com/naufalandika/designpattern/trial2/pb/certificate"
	"github.com/naufalandika/designpattern/trial2/pb/exercise"
)

type Service interface {
	SubmitPostTest(context.Context, *SubmitPostTestRequest) (*SubmitPostTestResponse, error)
}

func NewService(accomplishment *accomplishment.Accomplishment, certificate *certificate.Certificate, exercise *exercise.Exercise) Service {
	return &service{
		accomplishment: accomplishment,
		certificate:    certificate,
		exercise:       exercise,
	}
}

type service struct {
	accomplishment *accomplishment.Accomplishment
	certificate    *certificate.Certificate
	exercise       *exercise.Exercise
}

func (s *service) SubmitPostTest(ctx context.Context, req *SubmitPostTestRequest) (*SubmitPostTestResponse, error) {
	if err := s.validateSubmitPostTestProgress(ctx, req); err != nil {
		return nil, err
	}

	score := s.exercise.Submit()

	if score > 60 {
		s.generateCertificate()
	}

	return nil, nil
}

func (s *service) validateSubmitPostTestProgress(ctx context.Context, req *SubmitPostTestRequest) error {
	progress := s.accomplishment.GetUserProgress()

	if progress < 100 {
		return errors.New("user progress < 100")
	}

	return nil
}

func (s *service) generateCertificate() {
	s.certificate.Generate()
}
