package getquestionsprocessor

import (
	"context"
	"fmt"
)

type BaseProcessor struct{}

func (p *BaseProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	fmt.Println("base validation: progress = 100%")
	return nil
}
