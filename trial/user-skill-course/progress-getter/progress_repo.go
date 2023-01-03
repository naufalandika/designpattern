package progressgetter

import "context"

type IProgessRepo interface {
	GetProgress(context.Context, *RepoGetProgressRequest) (*RepoGetProgressResponse, error)
}

func NewProgressRepo() IProgessRepo {
	return &progressRepo{}
}

type progressRepo struct{}

func (r *progressRepo) GetProgress(ctx context.Context, req *RepoGetProgressRequest) (*RepoGetProgressResponse, error) {
	return &RepoGetProgressResponse{
		Progress: 80,
	}, nil
}
