package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
)

type nonprakerjaOnlineProcessor struct {
	accomplishment *accomplishment.Accomplishment
}

func (p *nonprakerjaOnlineProcessor) validateGetQuestionsProgress(ctx context.Context, req *validateGetQuestionsProgressRequest) error {
	progress := p.accomplishment.GetUserProgress()

	if progress < 100 {
		return errors.New("user progress < 100")
	}

	return nil
}
