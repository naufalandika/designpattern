package service

import "fmt"

type IService interface {
	GetTestSetQuestion()
	SubmitPostTest()
	postTestValidation()
	postTestPostSubmission()
}

func New(i IService) IService {
	if i == nil {
		return &Service{&Service{}}
	}

	return &Service{i}
}

type Service struct {
	iService IService
}

func (s *Service) GetTestSetQuestion() {
	s.iService.postTestValidation()

	fmt.Println("get questions")

	fmt.Println()
}

func (s *Service) SubmitPostTest() {
	s.iService.postTestValidation()

	fmt.Println("submit post test")

	s.iService.postTestPostSubmission()

	fmt.Println()
}

func (s *Service) postTestValidation() {
	fmt.Println("validasi: min progress 100%, submit pretest")
}

func (s *Service) postTestPostSubmission() {
	fmt.Println("generate sertifikat")
}
