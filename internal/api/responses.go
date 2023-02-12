package api

import (
	"encoding/json"
	"net/http"
)

func RenderJSONResponse(w http.ResponseWriter, v interface{}, code int) error {
	js, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(js)

	return nil
}
