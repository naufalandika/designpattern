package service

import "fmt"

type IService interface {
	GetQuestions()
	SubmitPostTest()
	postTestValidation()
	postTestPostSubmission()
}

func New(ct string) IService {
	var i IService
	switch ct {
	case "webinar":
		i = NewWebinarService()
	default:
		i = nil
	}

	if i == nil {
		return &Service{&Service{}}
	}

	return &Service{i}
}

type Service struct {
	iService IService
}

func (s *Service) GetQuestions() {
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
