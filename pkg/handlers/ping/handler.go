package ping

import (
	"bmstuInformaticsTechnologies/pkg/logging"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const (
	URL    = "/ping"
	METHOD = "GET"
)

type Handler struct {
	Logger logging.LoggerInterface
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc(URL, h.Ping).Methods(METHOD)
}

func (h *Handler) Ping(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"message": "pong"}`)
}
