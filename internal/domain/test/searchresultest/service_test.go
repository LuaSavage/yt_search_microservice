package searchresult_test

import (
	"context"
	"fmt"
	"testing"

	searchresult "github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	searchMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/searchresult"
	videoMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/video"
	apiMocks "github.com/LuaSavage/yt_search_microservice/pkg/mocks/ytsearch"
	ytVideoMocks "github.com/LuaSavage/yt_search_microservice/pkg/mocks/ytvideo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t testing.TB) (*searchMocks.Storage, *videoMocks.Service, *apiMocks.Client, searchresult.Service) {
	storage := searchMocks.NewStorage(t)
	videoService := videoMocks.NewService(t)

	api := apiMocks.NewClient(t)
	videoStreamApi := ytVideoMocks.NewClient(t)

	dto := &searchresult.NewServiceDTO{
		SearchApi:     api,
		Storage:       storage,
		VideoService:  videoService,
		YtVideoClient: videoStreamApi,
	}

	return storage, videoService, api, searchresult.NewService(dto)
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
		/*
			// loading predefined videos data
			videos := []video.Video{}

			buf, err := ioutil.ReadFile("../../searchresult/searchresult_testdata/data.yaml")
			require.NoError(t, err)

			err = yaml.Unmarshal([]byte(buf), &videos)
			require.NoError(t, err)

			// predefined search result dto
			videoIDs := []string{}

			for _, value := range videos {
				videoIDs = append(videoIDs, value.Id)
			}
		*/
		/*testSearchResultDTO := searchresult.StoreSearchResultDTO{
			Query:  "test quary",
			Videos: videoIDs,
		}*/

		ctx := context.TODO()

		for i := 0; i < len(testVideos); i++ {
			videoService.On("GetVideoByID", mock.AnythingOfType("*context.emptyCtx"), testVideos[i].Id).
				Return(&testVideos[i], nil).Once()
		}

		storage.On("GetSearchResultByQuary", mock.AnythingOfType("*context.emptyCtx"), mock.AnythingOfType("string")).
			Return(&testSearchResultDTO, nil).Once()

		//predefine search result
		/*testSearchResult := searchresult.SearchResult{
			Query:  testSearchResultDTO.Query,
			Videos: videos,
		}*/

		retrivedSearchResult, err := service.GetSearchResultByQuary(ctx, testSearchResult.Query)
		assert.NoError(t, err, "GetSearchResultByQuary ought to unmistakably works")
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
