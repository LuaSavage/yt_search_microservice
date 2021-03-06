package searchresult

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	searchURL = "/search"
)

type Handler interface {
	Register(router *httprouter.Router)
	Search(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params)
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
	router.GET(searchURL, h.Search)
}

func (h *handler) Search(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {

	query := request.URL.Query().Get("query")

	if len(query) > 0 {
		responseWriter.Header().Set("Content-Type", "application/json")
		searchResult, err := h.service.Search(request.Context(), query)

		if err == nil {
			responseWriter.WriteHeader(http.StatusOK)
			marshaled, _ := json.Marshal(searchResult)
			responseWriter.Write([]byte(marshaled))
		} else {
			responseWriter.WriteHeader(http.StatusNotFound)
			log.Fatal(err)
		}

	} else {
		responseWriter.WriteHeader(http.StatusNoContent)
	}
}
