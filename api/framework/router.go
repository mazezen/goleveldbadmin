package framework

import (
	"net/http"
)

var Router *RouterHandler = new(RouterHandler)

type RouterHandler struct{}

var mux = make(map[string]func(http.ResponseWriter, *http.Request))

func (p *RouterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if fun, ok := mux[r.URL.Path]; ok {
		fun(w, r)
		return
	}
	http.Error(w, "error URL: "+r.URL.Path, http.StatusBadRequest)
}

func (p *RouterHandler) Router(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	mux[pattern] = handler
}
