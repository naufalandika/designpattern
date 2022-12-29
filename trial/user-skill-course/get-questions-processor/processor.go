package getquestionsprocessor

import "context"

type GetQuestionsProcessor interface {
	Validate(context.Context, *ValidateRequest) error
}

func GetGetQuestionsProcessor(ctx context.Context, req *GetGetQuestionsProcessorRequest) GetQuestionsProcessor {
	return nil
}
