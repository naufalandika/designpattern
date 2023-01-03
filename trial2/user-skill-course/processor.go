package userskillcourse

import (
	"context"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
)

type validateGetQuestionsProgressRequest struct{}

type iProcessor interface {
	validateGetQuestionsProgress(context.Context, *validateGetQuestionsProgressRequest) error
}

type getProcessorRequest struct {
	userType   string
	courseType string
}

func getProcessor(req *getProcessorRequest) iProcessor {
	accomplishment := accomplishment.New()

	if req.userType == "prakerja" && req.courseType == "webinar" {
		return &prakerjaWebinarProcessor{accomplishment: accomplishment}
	}

	return &nonprakerjaOnlineProcessor{accomplishment: accomplishment}
}
