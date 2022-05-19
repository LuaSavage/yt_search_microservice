package searchresulttest

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	searchresult "github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	searchMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/searchmocks"
	videoMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/videomocks"
	apiMocks "github.com/LuaSavage/yt_search_microservice/pkg/mocks/ytsearchmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gopkg.in/yaml.v3"
)

func setup(t testing.TB) (*searchMocks.Storage, *videoMocks.Service, *apiMocks.Service, searchresult.Service) {
	storage := searchMocks.NewStorage(t)
	videoService := videoMocks.NewService(t)
	api := apiMocks.NewService(t)

	return storage, videoService, api, searchresult.NewService(storage, api, videoService)
}

// case of unexisting search results in cache
func TestGetSearchResultByQuaryERR(t *testing.T) {
	t.Run("error with cached search result", func(t *testing.T) {
		//storage, videoService, api, service := setup(t)
		storage, _, _, service := setup(t)

		errString := fmt.Errorf("a storage seems to be empty")
		ctx := context.TODO()
		storage.On("GetSearchResultByQuary", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).
			Return(nil, errString).Once()

		_, err := service.GetSearchResultByQuary(ctx, "test request")

		assert.Error(t, err)
	})
}

func TestGetSearchResultByQuaryOK(t *testing.T) {
	t.Run("error with cached search result", func(t *testing.T) {
		//storage, videoService, api, service := setup(t)
		storage, videoService, _, service := setup(t)

		// loading predefined videos data
		videos := []video.Video{}

		buf, err := ioutil.ReadFile("../../searchresult/searchresult_test_data/data.yaml")
		assert.Nil(t, err)

		err = yaml.Unmarshal([]byte(buf), &videos)
		assert.Nil(t, err)

		// predefined search result dto
		videoIDs := []string{}

		for _, value := range videos {
			videoIDs = append(videoIDs, value.Id)
		}

		testSearchResultDTO := searchresult.StoreSearchResultDTO{
			Query:  "test quary",
			Videos: videoIDs,
		}

		ctx := context.TODO()

		for i := 0; i < len(videos); i++ {
			videoService.On("GetVideoByID", mock.AnythingOfType("*context.emptyCtx"), videos[i].Id).
				Return(&videos[i], nil).Once()
		}

		storage.On("GetSearchResultByQuary", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).
			Return(&testSearchResultDTO, nil).Once()

		//predefine search result
		testSearchResult := searchresult.SearchResult{
			Query:  testSearchResultDTO.Query,
			Videos: videos,
		}

		retrivedSearchResult, err := service.GetSearchResultByQuary(ctx, testSearchResult.Query)
		assert.Nil(t, err, "GetSearchResultByQuary ought to unmistakably works")
		assert.Equal(t, len(testSearchResult.Videos), len(retrivedSearchResult.Videos))
		assert.ElementsMatch(t, testSearchResult.Videos, retrivedSearchResult.Videos)
	})
}

/*
func TestSearch(t *testing.T) {
	t.Run("error with cached search result", func(t *testing.T) {

	})
}
*/
