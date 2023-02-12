package api

import (
	"bmstuInformaticsTechnologies/pkg/handlers/notfound"
	"errors"
	"net/http"
)

type appHandler func(http.ResponseWriter, *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appErr *AppError
		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					notfound.NotFoundFunc(w, r)
					return
				}

				err := err.(*AppError)
				w.WriteHeader(err.HttpCode)
				w.Write(err.Marshal())
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(SystemError(err.Error()).Marshal())
		}
	}
}
