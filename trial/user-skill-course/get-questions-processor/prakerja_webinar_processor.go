package getquestionsprocessor

import (
	"context"
	"errors"

	progressgetter "github.com/naufalandika/designpattern/trial/user-skill-course/progress-getter"
)

type PrakerjaWebinarProcessor struct {
	progressGetter progressgetter.ProgressGetter
}

func (p *PrakerjaWebinarProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	userProgress, err := p.progressGetter.GetProgress(ctx, &progressgetter.GetProgressRequest{})
	if err != nil {
		return err
	}

	if userProgress.Progress < 80 {
		return errors.New("user progress < 80%")
	}

	return nil
}
