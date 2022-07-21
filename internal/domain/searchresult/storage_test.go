package searchresult

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	miniredis "github.com/alicebob/miniredis"
	redismock "github.com/elliotchance/redismock/v8"
	redis "github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

// newTestRedis returns a redis.Cmdable.
func newTestRedis() (*redis.Client, *redismock.ClientMock) {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client, redismock.NewNiceMock(client)
}

// initialise test object

var (
	ctx        context.Context = context.TODO()
	expiration int             = 100
)

func TestGetSearchResultByQuaryError(t *testing.T) {
	t.Run("error due to empty search result object response", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		searchResultStorage := NewStorage(redisClient, expiration)

		result, err := searchResultStorage.GetSearchResultByQuary(ctx, testSearchResultDTO.Query)
		assert.Error(t, err, fmt.Sprintf("retrived search result looks like %+v\n", result))
	})
}

// Writing SearchResultDTO and then trying to gather it
func TestGetSearchResultByQuaryOK(t *testing.T) {
	t.Run("error due to test search result dto doesnt received", func(t *testing.T) {
		redisClient, redisMock := newTestRedis()
		searchResultStorage := NewStorage(redisClient, expiration)

		redisMock.HSet(ctx, "search_result:"+testSearchResultDTO.Query, "query", testSearchResultDTO.Query)
		data, _ := json.Marshal(testSearchResultDTO.Videos)
		redisMock.HSet(ctx, "search_result:"+testSearchResultDTO.Query, "videos", data)

		retrivenResult, err := searchResultStorage.GetSearchResultByQuary(ctx, testSearchResultDTO.Query)
		assert.Equal(t, testSearchResultDTO, *retrivenResult, err)
	})
}

func SearchResultDtoToCrippledOrigin(dto StoreSearchResultDTO) SearchResult {
	innerVideos := []video.Video{}

	for _, id := range dto.Videos {
		innerVideos = append(innerVideos, video.Video{Id: id})
	}

	return SearchResult{
		Query:  dto.Query,
		Videos: innerVideos,
	}
}

func TestCreateSearchResultOK(t *testing.T) {
	t.Run("error due to the video  dosn't properly created", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		searchResultStorage := NewStorage(redisClient, expiration)

		_, err := searchResultStorage.GetSearchResultByQuary(ctx, testSearchResultDTO.Query)
		assert.Error(t, err)

		testSearchResult := SearchResultDtoToCrippledOrigin(testSearchResultDTO)

		err = searchResultStorage.CreateSearchResult(ctx, &testSearchResult)
		assert.NoError(t, err)

		// If it succesfully written then it has to be same as origin
		retrivenSearchResult, err := searchResultStorage.GetSearchResultByQuary(ctx, testSearchResult.Query)
		assert.Equal(t, testSearchResultDTO, *retrivenSearchResult, err)
	})
}

// prevention of double write issue
func TestCreateSearchResultError(t *testing.T) {
	t.Run("error due to double storring video up issue", func(t *testing.T) {
		redisClient, _ := newTestRedis()
		searchResultStorage := NewStorage(redisClient, expiration)
		err := searchResultStorage.CreateSearchResult(ctx, &testSearchResult)
		assert.NoError(t, err)
		err = searchResultStorage.CreateSearchResult(ctx, &testSearchResult)
		assert.Error(t, err)
	})
}
