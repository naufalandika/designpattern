package getquestionsprocessor

import (
	"context"
	"errors"
)

type PrakerjaOnlineGetQuestionsProcessor struct{}

func (p *PrakerjaOnlineGetQuestionsProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	return errors.New("user progress < 100%")
}
