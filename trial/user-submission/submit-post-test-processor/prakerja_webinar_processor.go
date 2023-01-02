package submitposttestprocessor

import (
	"context"
	"fmt"
)

type PrakerjaWebinarProcessor struct{}

func (p *PrakerjaWebinarProcessor) Validate(context.Context, *ValidateRequest) error {
	fmt.Println("pretest > 0 && progress > 80%")
	return nil
}

func (p *PrakerjaWebinarProcessor) PostSubmission(context.Context, *PostSubmissionRequest) error {
	fmt.Println("if score >= 60 maka update uac")
	return nil
}
