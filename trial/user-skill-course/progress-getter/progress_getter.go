package progressgetter

import "context"

type ProgressGetter interface {
	GetProgress(context.Context, *GetProgressRequest) (*GetProgressResponse, error)
}

func New() ProgressGetter {
	return nil
}
