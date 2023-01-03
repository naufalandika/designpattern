package userskillcourse

import (
	"context"
	"fmt"
)

type nonPrakerjaOnlineProcessor struct{}

func (p *nonPrakerjaOnlineProcessor) validate(ctx context.Context, req *getQuestionsProcessorValidateRequest) error {
	fmt.Println("non prakerja online: user progress = 100%")
	return nil
}
