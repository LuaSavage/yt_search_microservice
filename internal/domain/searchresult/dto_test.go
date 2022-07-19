package searchresult

import (
	"io/ioutil"
	"testing"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestNewStoreSearchResultDTO(t *testing.T) {

	t.Run("error of assembling SearchResultDTO", func(t *testing.T) {
		// loading predefined videos data
		videos := []video.Video{}

		buf, err := ioutil.ReadFile("../searchresult_testdata/data.yaml")
		require.NoError(t, err)

		err = yaml.Unmarshal([]byte(buf), &videos)
		require.NoError(t, err)

		testSearchResult := &SearchResult{
			Query:  "some question",
			Videos: videos,
		}

		ids := []string{}
		for _, testVideo := range videos {
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
