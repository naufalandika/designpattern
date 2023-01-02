package usersubmission

import "context"

type Service interface {
	SubmitPostTest(context.Context, *SubmitPostTestRequest) (*SubmitPostTestResponse, error)
}

func NewService() Service {
	return nil
}
