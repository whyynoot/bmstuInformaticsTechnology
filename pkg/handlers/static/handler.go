package static

import (
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	URL = "/static/"
	DIR = "./static/"
)

type Handler struct {
	Logger logging.LoggerInterface
}

func (h *Handler) Register(router *mux.Router) {
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		http.FileServer(http.Dir("./static/"))))
}
