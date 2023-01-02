package submitposttestprocessor

import "context"

type Processor interface {
	Validate(context.Context, *ValidateRequest) error
	PostSubmission(context.Context, *PostSubmissionRequest) error
}

func GetProcessor(ctx context.Context, req *GetProcessorRequest) Processor {
	return &BaseProcessor{}
}
