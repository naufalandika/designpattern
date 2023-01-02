package usersubmission

import "context"

type Service interface {
	SubmitPostTest(context.Context, *SubmitPostTestRequest) (*SubmitPostTestResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) SubmitPostTest(ctx context.Context, req *SubmitPostTestRequest) (*SubmitPostTestResponse, error) {
	return nil, nil
}
