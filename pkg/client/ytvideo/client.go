package ytvideo

import (
	"context"
	"io"
	"net/http"

	youtube "github.com/kkdai/youtube/v2"
)

type Client interface {
	GetStream(video *youtube.Video, format *youtube.Format) (io.ReadCloser, int64, error)
	GetStreamContext(ctx context.Context, video *youtube.Video, format *youtube.Format) (io.ReadCloser, int64, error)
	GetStreamURL(video *youtube.Video, format *youtube.Format) (string, error)
	GetStreamURLContext(ctx context.Context, video *youtube.Video, format *youtube.Format) (string, error)
	GetVideo(url string) (*youtube.Video, error)
	GetVideoContext(ctx context.Context, url string) (*youtube.Video, error)
}

func NewClient(httpClient *http.Client) Client {
	ytClient := youtube.Client{
		Debug:      false,
		HTTPClient: httpClient,
	}

	return &ytClient
}
