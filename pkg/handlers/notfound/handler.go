package notfound

import (
	"net/http"
)

const (
	URL = "/static/404.html"
)

func NotFoundHandler() http.Handler {
	return http.HandlerFunc(NotFoundFunc)
}

func NotFoundFunc(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, URL, http.StatusPermanentRedirect)
}
