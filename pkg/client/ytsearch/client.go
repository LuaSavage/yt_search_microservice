package ytsearchapi

import (
	YTSearch "github.com/AnjanaMadu/YTSearch"
)

//mockery --name=Client --filename=client.go --output=../../mocks/ytsearch/ --outpkg=ytsearchmocks
type Client interface {
	Search(query string) (result []*ResultDTO, err error)
}

type client struct{}

func NewService() Client {
	return &client{}
}

func (c *client) Search(query string) (result []*ResultDTO, err error) {
	results, err := YTSearch.Search(query)

	if err != nil {
		return nil, err
	}

	for _, res := range results {
		result = append(result, &ResultDTO{
			Title:       res.Title,
			Id:          res.VideoId,
			PublishTime: res.PublishTime,
			Channel:     res.Channel,
			//ChannelId:   res.ChannelId,
			Views: res.Views,
			//Duration:    res.Duration,
			Thumbnail: res.Thumbnail,
		})
	}

	return result, nil
}
