package getquestionsprocessor

import "context"

type PrakerjaOnlineGetQuestionsProcessor struct{}

func (p *PrakerjaOnlineGetQuestionsProcessor) Validate(ctx context.Context, req *ValidateRequest) error {
	return nil
}
