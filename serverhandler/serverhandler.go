package serverhandler

import (
	"gopkg.in/mgo.v2"
	"net/http"
)

type ServerHandler struct {
	preFilters, postFilters []func(http.ResponseWriter, *http.Request)
}

func (sh *ServerHandler) AddPreFilters(f ...func(http.ResponseWriter, *http.Request)) {
	sh.preFilters = append(sh.preFilters,f...)
}

func (sh *ServerHandler) AddPostFilters(f ...func(http.ResponseWriter, *http.Request)) {
	sh.postFilters = append(sh.postFilters,f...)
}

func (sh ServerHandler) Func(db *mgo.Database, f func(*mgo.Database)func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, pre := range sh.preFilters {
			pre(w, r)
		}
		f(db)(w, r)
		for _,post := range sh.postFilters {
			post(w, r)
		}
	}
}