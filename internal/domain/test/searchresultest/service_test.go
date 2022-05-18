package searchresulttest

import (
	"context"
	"fmt"
	"testing"

	searchresult "github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	searchMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/searchmocks"
	videoMocks "github.com/LuaSavage/yt_search_microservice/internal/mocks/videomocks"
	apiMocks "github.com/LuaSavage/yt_search_microservice/pkg/mocks/ytsearchmocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(t testing.TB) (*searchMocks.Storage, *videoMocks.Service, *apiMocks.Service, searchresult.Service) {
	storage := searchMocks.NewStorage(t)
	videoService := videoMocks.NewService(t)
	api := apiMocks.NewService(t)

	return storage, videoService, api, searchresult.NewService(storage, api, videoService)
}

func TestGetSearchResultByQuary(t *testing.T) {
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

/*
func TestSearch(t *testing.T) {
	t.Run("error with cached search result", func(t *testing.T) {

	})
}*/
