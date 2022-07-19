package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/LuaSavage/yt_search_microservice/internal/config"
	"github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	"github.com/LuaSavage/yt_search_microservice/internal/domain/video"
	"github.com/LuaSavage/yt_search_microservice/internal/domain/videostream"
	"github.com/LuaSavage/yt_search_microservice/pkg/client/cache"
	ytsearchapi "github.com/LuaSavage/yt_search_microservice/pkg/client/ytsearch"
	"github.com/LuaSavage/yt_search_microservice/pkg/client/ytvideo"
	"github.com/alicebob/miniredis"
	"github.com/julienschmidt/httprouter"
)

func main() {
	cfg, err := config.GetConfig("config.yml")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Create http router")
	router := httprouter.New()

	log.Println("Create search row handler")

	//mock redis
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	fmt.Println(mr.Addr())

	redisHost := fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port)
	cacheClient, err := cache.NewClient(redisHost, cfg.Redis.Password, cfg.Redis.Database)

	if err != nil {
		log.Fatal(err)
	}

	// initialising storages
	searchResultStorage := searchresult.NewStorage(cacheClient, cfg.ObjectTTL.SearchResult)
	videoStorage := video.NewStorage(cacheClient, cfg.ObjectTTL.Video)
	videoStreamPoolStorage := videostream.NewStorage(cacheClient, cfg.ObjectTTL.VideoStreamPool)

	// initialising services
	videoService := video.NewService(videoStorage)

	httClient := &http.Client{}
	ytVideoClient := ytvideo.NewClient(httClient)
	searchApi := ytsearchapi.NewClient()

	searchService := searchresult.NewService(&searchresult.NewServiceDTO{
		SearchApi:     searchApi,
		Storage:       searchResultStorage,
		VideoService:  videoService,
		YtVideoClient: ytVideoClient,
	})

	videoStreamPoolService := videostream.NewService(videoStreamPoolStorage, ytVideoClient)

	// initialising handlers
	searchHandler := searchresult.NewHandler(searchService)
	searchHandler.Register(router)

	videoStreamService := videostream.NewHandler(videoStreamPoolService)
	videoStreamService.Register(router)

	start(router)
}

func start(router *httprouter.Router) {

	//log.Infof("bind application to host: %s and port: %s", cfg.Listen.BindIP, cfg.Listen.Port)

	var err error

	//listener, err = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", "127.0.0.1", "321"))

	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err = server.Serve(listener)

	if err != nil {
		log.Fatal(err)
	}

}
