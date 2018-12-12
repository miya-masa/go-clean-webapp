package web

import (
	"encoding/json"
	"net/http"
)

func jsonView(w http.ResponseWriter, code int, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
	w.WriteHeader(code)
	return nil
}
