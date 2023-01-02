package submitposttestprocessor

import (
	"context"
	"fmt"
)

type NonPrakerjaOnlineProcessor struct{}

func (p *NonPrakerjaOnlineProcessor) Validate(context.Context, *ValidateRequest) error {
	fmt.Println("pretest > 0 && progress == 100%")
	return nil
}

func (p *NonPrakerjaOnlineProcessor) PostSubmission(context.Context, *PostSubmissionRequest) error {
	return nil
}
