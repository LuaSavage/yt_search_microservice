package videostream

type VideoStream struct {
	VideoId string `json:"videoId" redis:"videoId" yaml:"video_id"`
	Quality string `json:"quality" redis:"qality" yaml:"quality"`
	Url     string `json:"url" redis:"url" yaml:"url"`
}

type VideoStreamPool struct {
	VideoId string        `json:"videoId" redis:"videoId" yaml:"video_id"`
	Streams []VideoStream `json:"streams" redis:"streams" yaml:"streams"`
}
