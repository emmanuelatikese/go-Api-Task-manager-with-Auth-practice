package apiUtils

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(content interface{}, w http.ResponseWriter, code int) {
	jsonStr, err := json.Marshal(content)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonStr)
}