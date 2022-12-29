package service

import "fmt"

type WebinarService struct{}

func (s *WebinarService) GetQuestions() {
	fmt.Println("user progress must be > 80%")
	fmt.Println("this is the questions")
}
