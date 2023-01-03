package progress

import (
	"context"

	"github.com/naufalandika/designpattern/trial2/pb/accomplishment"
)

type GetRequest struct{}

type GetResponse struct {
	UserProgress int32
}

type Progress interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func New(accomplishment *accomplishment.Accomplishment) Progress {
	return &progress{accomplishment}
}

type progress struct {
	accomplishment *accomplishment.Accomplishment
}

func (p *progress) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	progress := p.accomplishment.GetUserProgress()
	return &GetResponse{
		UserProgress: progress,
	}, nil
}
