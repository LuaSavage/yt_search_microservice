package video

type Video struct {
	Title       string `json:"title"`
	Id          string `json:"videoId"`
	PublishTime string `json:"publishTime"`
	Channel     string `json:"channel"`
	Views       string `json:"views"`
	Thumbnail   string `json:"thumbnail"`
}
