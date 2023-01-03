package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/progress"
)

type nonprakerjaOnlineProcessor struct {
	progress progress.Progress
}

func (p *nonprakerjaOnlineProcessor) validateGetQuestionsProgress(ctx context.Context, req *validateGetQuestionsProgressRequest) error {
	progress, err := p.progress.Get(ctx, &progress.GetRequest{})
	if err != nil {
		return err
	}

	if progress.UserProgress < 100 {
		return errors.New("user progress < 80")
	}

	return nil
}
