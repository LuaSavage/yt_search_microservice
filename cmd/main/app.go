package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/LuaSavage/yt_search_microservice/internal/domain/searchresult"
	"github.com/julienschmidt/httprouter"
)

func main() {
	/*cfg, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
	}*/

	log.Println("Create http router")
	router := httprouter.New()

	log.Println("Create search row handler")

	searchService := searchresult.NewService(&searchresult.NewServiceDTO{
		SearchApi:     nil,
		Storage:       nil,
		VideoService:  nil,
		YtVideoClient: nil,
	})

	handler := searchresult.NewHandler(searchService)
	handler.Register(router)

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
