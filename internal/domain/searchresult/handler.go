package searchresult

import (
	"context"
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
	responseWriter.WriteHeader(200)
	queryValues := request.URL.Query()

	query := queryValues.Get("query")

	if len(query) > 0 {
		ctx := context.TODO()
		responseWriter.WriteHeader(200)
		h.service.GetSearchResultByQuary(ctx, query)
	} else {
		responseWriter.WriteHeader(204)
	}

	responseWriter.Write([]byte("Search result " + queryValues.Get("query")))
}
