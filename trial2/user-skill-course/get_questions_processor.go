package userskillcourse

import "context"

type getQuestionsProcessorValidateRequest struct{}

type getQuestionsProcessor interface {
	validate(context.Context, *getQuestionsProcessorValidateRequest) error
}

func newGetQuestionsProcessor() getQuestionsProcessor {
	return &nonPrakerjaOnlineProcessor{}
}
