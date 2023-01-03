package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/progress"
)

type prakerjaWebinarProcessor struct {
	progress progress.Progress
}

func (p *prakerjaWebinarProcessor) validateGetQuestionsProgress(ctx context.Context, req *validateGetQuestionsProgressRequest) error {
	progress, err := p.progress.Get(ctx, &progress.GetRequest{})
	if err != nil {
		return err
	}

	if progress.UserProgress < 80 {
		return errors.New("user progress < 80")
	}

	return nil
}
