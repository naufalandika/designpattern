package getquestionsprocessor

import (
	"context"
	"fmt"
)

type PrakerjaWebinarGetQuestionsProcessor struct{}

func (p *PrakerjaWebinarGetQuestionsProcessor) Validate(context.Context, *ValidateRequest) error {
	fmt.Println("webinar get question validation: progress > 80%")
	return nil
}
