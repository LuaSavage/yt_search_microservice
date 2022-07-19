package videostream

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	searchURL = "/video"
)

type Handler interface {
	Register(router *httprouter.Router)
	GetVideoStreamPool(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	log.Println("Register handler " + searchURL)
	router.GET(searchURL, h.GetVideoStreamPool)
}

func (h *handler) GetVideoStreamPool(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	id := request.URL.Query().Get("id")

	if len(id) > 0 {
		responseWriter.Header().Set("Content-Type", "application/json")

		videoStreamPool, err := h.service.Get(request.Context(), id)

		if err == nil {
			responseWriter.WriteHeader(http.StatusOK)
			marshaled, _ := json.Marshal(videoStreamPool)
			responseWriter.Write([]byte(marshaled))
		} else {
			responseWriter.WriteHeader(http.StatusNotFound)
			log.Fatal(err)
		}

	} else {
		responseWriter.WriteHeader(http.StatusNoContent)
	}
}
