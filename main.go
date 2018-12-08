package main

import (
	"encoding/json"
	"fmt"
	"github.com/gaoyuyue/feed-backend/filter"
	"github.com/gaoyuyue/feed-backend/handler"
	"github.com/gaoyuyue/feed-backend/serverhandler"
	"gopkg.in/mgo.v2"
	"log"
	"net/http"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db := session.DB("feed")

	server := &serverhandler.ServerHandler{}
	server.AddPreFilters(filter.Cors)

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/add_one", server.Func(db, handler.AddFeed))
	http.HandleFunc("/list", server.Func(db, handler.ListFeed))
	http.HandleFunc("/delete_one", server.Func(db, handler.DeleteFeed))
	http.HandleFunc("/update_one", server.Func(db, handler.UpdateFeed))
	log.Fatal(http.ListenAndServe("0.0.0.0:80", nil))
	//log.Fatal(http.ListenAndServeTLS("0.0.0.0:80", "ca.crt","ca.key", nil))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	data, _ := json.Marshal(map[string]interface{}{
		"name":       "小年糕",
		"age":        5,
		"url":        url,
		"method":     r.Method,
		"proto":      r.Proto,
		"header":     r.Header,
		"host":       r.Host,
		"form":       r.Form,
		"RemoteAddr": r.RemoteAddr,
	})
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, string(data))
}