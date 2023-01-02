package getquestionsprocessor

import (
	"context"
	"fmt"
)

type PrakerjaWebinarProcessor struct{}

func (p *PrakerjaWebinarProcessor) Validate(context.Context, *ValidateRequest) error {
	fmt.Println("webinar get question validation: progress > 80%")
	return nil
}
