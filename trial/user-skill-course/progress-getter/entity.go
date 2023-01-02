package progressgetter

type GetProgressRequest struct{}

type GetProgressResponse struct {
	Progress float64
}

type RepoGetProgressRequest struct{}

type RepoGetProgressResponse struct {
	Progress float64
}
