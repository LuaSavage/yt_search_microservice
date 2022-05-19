package video

type Video struct {
	Title       string `json:"title" redis:"title" yaml:"title"`
	Id          string `json:"videoId" redis:"videoId" yaml:"video_id"`
	PublishTime string `json:"publishTime" redis:"publishTime" yaml:"publish_time"`
	Channel     string `json:"channel" redis:"channel" yaml:"channel"`
	Views       string `json:"views" redis:"views" yaml:"views"`
	Thumbnail   string `json:"thumbnail" redis:"thumbnail" yaml:"thumbnail"`
}
