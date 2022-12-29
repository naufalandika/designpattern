package service

import "fmt"

type IService interface {
	GetQuestions()
}

func New(ct string) IService {
	if ct == "webinar" {
		return &WebinarService{}
	}

	return &Service{}
}

type Service struct{}

func (s *Service) GetQuestions() {
	fmt.Println("user progress must be 100%")
	fmt.Println("this is the questions")
}
