package ytsearchapi

import (
	YTSearch "github.com/AnjanaMadu/YTSearch"
)

type Service interface {
	Search(query string) (result []ResultDTO, err error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Search(query string) (result []ResultDTO, err error) {
	results, err := YTSearch.Search(query)

	if err != nil {
		return nil, err
	}

	for _, res := range results {
		result = append(result, ResultDTO{
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
