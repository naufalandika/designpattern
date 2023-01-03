package submitposttestprocessor

import (
	"context"
	"fmt"
)

type PrakerjaOnlineProcessor struct{}

func (p *PrakerjaOnlineProcessor) Validate(context.Context, *ValidateRequest) error {
	fmt.Println("pretest > 0 && progress == 100%")
	return nil
}

func (p *PrakerjaOnlineProcessor) PostSubmission(context.Context, *PostSubmissionRequest) error {
	fmt.Println("generate coc")
	fmt.Println("update uac")

	fmt.Println("generate coe if eligible")
	fmt.Println("update uac")

	fmt.Println("call bukalapak reporting api")
	return nil
}
