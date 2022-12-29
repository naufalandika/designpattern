package main

import "github.com/naufalandika/designpattern/template_method/service"

func main() {
	courseType := "webinar"

	var i service.IService
	switch courseType {
	case "webinar":
		i = service.NewWebinarService()
	default:
		i = nil
	}

	svc := service.New(i)

	svc.GetTestSetQuestion()
	svc.SubmitPostTest()
}
