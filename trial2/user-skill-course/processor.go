package userskillcourse

import (
	"context"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
	"github.com/naufalandika/designpattern/trial2/progress"
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
	progress := progress.New(accomplishment)

	if req.userType == "prakerja" && req.courseType == "webinar" {
		return &prakerjaWebinarProcessor{progress: progress}
	}

	return &nonprakerjaOnlineProcessor{accomplishment: accomplishment}
}
