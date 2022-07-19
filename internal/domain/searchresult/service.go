package searchresult

import (
	"context"
	"log"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	ytsearch "github.com/LuaSavage/yt_search_microservice/pkg/client/ytsearch"
	ytvideo "github.com/LuaSavage/yt_search_microservice/pkg/client/ytvideo"
)

type Service interface {
	GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error)
	Search(ctx context.Context, query string) (*SearchResult, error)
}

type service struct {
	searchApi     ytsearch.Client
	storage       Storage
	videoService  video.Service
	ytVideoClient ytvideo.Client
}

func NewService(dto *NewServiceDTO) Service {
	return &service{
		searchApi:     dto.SearchApi,
		storage:       dto.Storage,
		videoService:  dto.VideoService,
		ytVideoClient: dto.YtVideoClient}
}

func (s *service) GetSearchResultByQuary(ctx context.Context, query string) (*SearchResult, error) {
	resultsDTO, err := s.storage.GetSearchResultByQuary(ctx, query)

	if err != nil {
		return nil, err
	}

	// now we gonna pull up from storage all of the videos by id
	results := SearchResult{
		Query:  query,
		Videos: []video.Video{},
	}

	for _, videoId := range resultsDTO.Videos {
		retrivenVideo, err := s.videoService.GetVideoByID(ctx, videoId)

		if err == nil {
			results.Videos = append(results.Videos, *retrivenVideo)
		} else {
			//case of unexisting video in cache
			//we better request it by iteslf
			ytVideo, err := s.ytVideoClient.GetVideoContext(ctx, videoId)

			if err == nil {
				obtainedVideo := video.Video{
					Title:       ytVideo.Title,
					Id:          videoId,
					PublishTime: ytVideo.PublishDate.Format("2006-01-02"),
					Channel:     ytVideo.Author,
					Views:       "-/-",
					Thumbnail:   ytVideo.Thumbnails[0].URL,
				}

				results.Videos = append(results.Videos, obtainedVideo)

				// at the same same time it reasonable to extract streams urls, but not now
			}
		}
	}

	return &results, nil
}

func (s *service) Search(ctx context.Context, query string) (*SearchResult, error) {

	// If is here ready result then returning them
	results, err := s.GetSearchResultByQuary(ctx, query)

	if err == nil {
		log.Printf("%s '%s' %s", "search result", query, "found in cache")
		return results, nil
	}

	result, err := s.searchApi.Search(query)

	if err != nil {
		log.Fatal(err)
	}

	// Searching for new results from lib's api
	var videoPool []video.Video

	for _, res := range result {
		videoPool = append(videoPool, video.Video(*res))
	}

	currentSearchResults := &SearchResult{
		Query:  query,
		Videos: videoPool}

	return currentSearchResults, nil
}
