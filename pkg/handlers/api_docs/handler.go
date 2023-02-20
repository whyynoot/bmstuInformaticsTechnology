package api_docs

import (
	_ "bmstuInformaticsTechnologies/docs"
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	URL    = "/documentation"
	METHOD = "GET"
)

type Handler struct {
	Logger logging.LoggerInterface
}

func (h *Handler) Register(router *mux.Router) {
	router.PathPrefix(URL).Handler(httpSwagger.WrapHandler).Methods("GET")
}
