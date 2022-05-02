package video

type Video struct {
	Id          string `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	PublishTime string `json:"publishTime"`
	Channel     string `json:"channel"`
	Views       string `json:"views"`
}
