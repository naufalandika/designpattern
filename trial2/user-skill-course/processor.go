package userskillcourse

import "context"

type validateGetQuestionsProgressRequest struct{}

type iProcessor interface {
	validateGetQuestionsProgress(context.Context, *validateGetQuestionsProgressRequest) error
}

type getProcessorRequest struct {
	userType   string
	courseType string
}

func getProcessor(req *getProcessorRequest) iProcessor {
	return &nonprakerjaOnlineProcessor{}
}
