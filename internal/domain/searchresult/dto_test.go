package searchresult

import (
	"testing"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	"github.com/stretchr/testify/assert"
)

var (
	testVideos []video.Video = []video.Video{
		{
			Title:       "Video 1",
			Id:          "=asdSFAASs",
			PublishTime: "01-02-2015",
			Channel:     "jabbascript",
			Views:       "1M",
			Thumbnail:   "http://someshit.d",
		},
		{
			Title:       "Video 2",
			Id:          "=asdSFfghSs",
			PublishTime: "28-08-1986",
			Channel:     "cobol",
			Views:       "1D",
			Thumbnail:   "http://deepshit.b",
		}}

	testSearchResult SearchResult = SearchResult{
		Query:  "some question",
		Videos: testVideos,
	}
)

func TestNewStoreSearchResultDTO(t *testing.T) {
	t.Run("error of assembling SearchResultDTO", func(t *testing.T) {

		ids := []string{}
		for _, testVideo := range testVideos {
			ids = append(ids, testVideo.Id)
		}

		// struct to verify
		testSearchResultDTO := StoreSearchResultDTO{
			Query:  testSearchResult.Query,
			Videos: ids,
		}

		retrived := NewStoreSearchResultDTO(testSearchResult)
		assert.Equal(t, testSearchResultDTO, retrived)
	})
}
