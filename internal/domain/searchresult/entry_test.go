package searchresult

import (
	"log"
	"os"
	"testing"

	video "github.com/LuaSavage/yt_search_microservice/internal/domain/video"
)

var (
	testSearchResultDTO StoreSearchResultDTO
	testSearchResult    SearchResult
	testVideos          []video.Video
)

// main entry point of all package tests
func TestMain(m *testing.M) {

	err := LoadFromYaml("./searchresult_testdata/search_result.yaml", &testSearchResult)
	if err != nil {
		log.Fatal(err)
	}

	err = LoadFromYaml("./searchresult_testdata/results.yaml", &testVideos)
	if err != nil {
		log.Fatal(err)
	}
	testSearchResult.Videos = testVideos

	err = LoadFromYaml("./searchresult_testdata/search_result_dto.yaml", &testSearchResultDTO)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
