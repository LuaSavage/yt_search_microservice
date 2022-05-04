package searchresult

type StoreSearchResultDTO struct {
	Query  string   `json:"query"`
	Videos []string `json:"videos"`
}

func NewStoreSearchResultDTO(searchResult SearchResult) StoreSearchResultDTO {
	var videoIds []string

	result := StoreSearchResultDTO{Query: searchResult.Query}

	for _, video := range searchResult.Videos {
		videoIds = append(videoIds, video.Id)
	}

	result.Videos = videoIds
	return result
}
