package video

type Video struct {
	Title       string `json:"title" redis:"title"`
	Id          string `json:"videoId" redis:"videoId"`
	PublishTime string `json:"publishTime" redis:"publishTime"`
	Channel     string `json:"channel" redis:"channel"`
	Views       string `json:"views" redis:"views"`
	Thumbnail   string `json:"thumbnail" redis:"thumbnail"`
}
