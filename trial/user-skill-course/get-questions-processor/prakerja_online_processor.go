package getquestionsprocessor

import (
	"context"
	"errors"

	progressgetter "github.com/naufalandika/designpattern/trial/user-skill-course/progress-getter"
)

type PrakerjaOnlineProcessor struct {
	ProgressGetter progressgetter.ProgressGetter
}

func (p *PrakerjaOnlineProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	userProgress, err := p.ProgressGetter.GetProgress(ctx, &progressgetter.GetProgressRequest{})
	if err != nil {
		return err
	}

	if userProgress.Progress < 100 {
		return errors.New("user progress < 100%")
	}

	return nil
}
