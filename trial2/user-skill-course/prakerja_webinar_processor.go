package userskillcourse

import (
	"context"
	"errors"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
)

type prakerjaWebinarProcessor struct {
	accomplishment *accomplishment.Accomplishment
}

func (p *prakerjaWebinarProcessor) validateGetQuestionsProgress(context.Context, *validateGetQuestionsProgressRequest) error {
	progress := p.accomplishment.GetUserProgress()

	if progress < 80 {
		return errors.New("user progress < 80")
	}

	return nil
}
