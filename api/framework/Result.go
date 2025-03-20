package framework

import (
	"encoding/json"
	"io"
	"net/http"
)

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Ok(w http.ResponseWriter, data string) {
	_, _ = io.WriteString(w, data)
}

func Fail(w http.ResponseWriter, err string) {
	http.Error(w, err, http.StatusBadRequest)
}

func JsonOk(w http.ResponseWriter, data *Result) {
	w.Header().Set("Content-Type", "application/json")
	b, _ := json.Marshal(data)
	_, _ = w.Write(b)
}
