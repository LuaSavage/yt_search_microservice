package searchresult

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStoreSearchResultDTO(t *testing.T) {

	t.Run("error of assembling SearchResultDTO", func(t *testing.T) {
		retrived := NewStoreSearchResultDTO(&testSearchResult)
		assert.Equal(t, testSearchResultDTO, *retrived)
	})
}
