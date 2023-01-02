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
	fmt.Println("generate coc")
	fmt.Println("update uac")

	fmt.Println("generate coe if eligible")
	fmt.Println("update uac")
	return nil
}
