package submitposttestprocessor

import (
	"context"
	"fmt"
)

type BaseProcessor struct{}

func (p *BaseProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	fmt.Println("pretest > 0 && progress = 100%")
	return nil
}

func (p *BaseProcessor) PostSubmission(context.Context, *PostSubmissionRequest) error {
	// have no post submission
	return nil
}
