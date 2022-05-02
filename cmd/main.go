package main

import (
	"fmt"
	"time"

	ytsearch "github.com/AnjanaMadu/YTSearch"
	youtube "github.com/kkdai/youtube/v2"
)

type VideDesc struct {
	OrderId int
	VideoID string
}

func SearchVideo(videoID string, orderID int, channel chan VideDesc, Client *youtube.Client) {
	if len(videoID) > 0 {

		_, err := Client.GetVideo(videoID)
		if err != nil {
			fmt.Println(err)
		}

		//fmt.Println(orderID, videoID)

		/*formats := video.Formats.WithAudioChannels() // only get videos with audio
		stream, _, err := client.GetStream(video, &formats[0])
		if err != nil {
			panic(err)
		}

		file, err := os.Create("video.mp4")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			panic(err)
		}*/
	}

	vd := VideDesc{orderID, videoID}

	channel <- vd
}

func main() {
	results, err := ytsearch.Search("jurrasik park ambience")

	start := time.Now()
	// some computation

	if err != nil {
		fmt.Println(err)
	}

	client := youtube.Client{}
	minedOreChan := make(chan VideDesc)

	for num, result := range results {
		go SearchVideo(result.VideoId, num, minedOreChan, &client)
		time.Sleep(500 * time.Millisecond)
	}

	for i := 0; i < len(results); i++ {
		go fmt.Println(<-minedOreChan)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

	<-time.After(time.Second * 1)
}
