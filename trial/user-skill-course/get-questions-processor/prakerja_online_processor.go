package getquestionsprocessor

import (
	"context"
	"errors"
)

type PrakerjaOnlineProcessor struct{}

func (p *PrakerjaOnlineProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	return errors.New("user progress < 100%")
}
