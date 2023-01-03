package progressgetter

import "context"

type IProgressGetter interface {
	GetProgress(context.Context, *GetProgressRequest) (*GetProgressResponse, error)
}

func NewProgressGetter(repo IProgessRepo) IProgressGetter {
	return &progressGetter{repo}
}

type progressGetter struct {
	repo IProgessRepo
}

func (g *progressGetter) GetProgress(ctx context.Context, req *GetProgressRequest) (*GetProgressResponse, error) {
	resp, err := g.repo.GetProgress(ctx, &RepoGetProgressRequest{})
	if err != nil {
		return nil, err
	}

	return &GetProgressResponse{
		Progress: resp.Progress,
	}, nil
}
