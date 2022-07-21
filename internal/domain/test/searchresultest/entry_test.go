package searchresult_test

import (
	"log"
	"os"
	"testing"

	searchresult "github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
)

var (
	testSearchResultDTO searchresult.StoreSearchResultDTO
	testSearchResult    searchresult.SearchResult
	testVideos          []video.Video
)

// main entry point of all package tests
func TestMain(m *testing.M) {

	err := LoadFromYaml("../../searchresult/searchresult_testdata/results.yaml", &testVideos)
	if err != nil {
		log.Fatal(err)
	}
	testSearchResult.Videos = testVideos

	err = LoadFromYaml("../../searchresult/searchresult_testdata/search_result.yaml", &testSearchResult)
	if err != nil {
		log.Fatal(err)
	}

	err = LoadFromYaml("../../searchresult/searchresult_testdata/search_result_dto.yaml", &testSearchResultDTO)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
