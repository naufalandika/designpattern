package service

import "fmt"

type WebinarService struct {
	Service
}

func NewWebinarService() IService {
	return &WebinarService{}
}

func (s *WebinarService) postTestValidation() {
	fmt.Println("validasi: min progress 80%")
}
