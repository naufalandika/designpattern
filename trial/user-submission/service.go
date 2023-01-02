package usersubmission

import (
	"context"
	"fmt"
)

type Service interface {
	SubmitPostTest(context.Context, *SubmitPostTestRequest) (*SubmitPostTestResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (s *service) SubmitPostTest(ctx context.Context, req *SubmitPostTestRequest) (*SubmitPostTestResponse, error) {
	fmt.Println("validation")

	fmt.Println("base logic submit post test")

	fmt.Println("post submission")

	return nil, nil
}
