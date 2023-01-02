package getquestionsprocessor

import (
	"context"
	"errors"

	progressgetter "github.com/naufalandika/designpattern/trial/user-skill-course/progress-getter"
)

type BaseProcessor struct {
	progressGetter progressgetter.IProgressGetter
}

func (p *BaseProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	userProgress, err := p.progressGetter.GetProgress(ctx, &progressgetter.GetProgressRequest{})
	if err != nil {
		return err
	}

	if userProgress.Progress < 100 {
		return errors.New("user progress < 100%")
	}

	return nil
}
