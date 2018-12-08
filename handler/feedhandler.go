package handler

import (
	"encoding/json"
	"github.com/gaoyuyue/feed-backend/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

func AddFeed(db *mgo.Database) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			values := r.Form
			title := values.Get("title")
			content := values.Get("content")
			feed := &entity.Feed{Id:bson.NewObjectId(),Title:title, Content:content, CreateTime:time.Now()}
			db.C("feeds").Insert(feed)
			if bs, err := json.Marshal(*feed); err == nil {
				w.Write(bs)
			}
			w.WriteHeader(200)
		} else if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(401)
		}
	}
}

func ListFeed(db *mgo.Database) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			result := make([]entity.Feed, 0)
			db.C("feeds").Find(nil).All(&result)
			bs, err := json.Marshal(result)
			if err != nil {
				w.WriteHeader(500)
			} else {
				w.Write(bs)
			}
		} else if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(401)
		}
	}
}


func DeleteFeed(db *mgo.Database) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			id := r.URL.Query().Get("id")
			db.C("feeds").Remove(bson.M{"_id": bson.ObjectIdHex(id)})
			w.WriteHeader(200)
		} else if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(401)
		}
	}
}

func UpdateFeed(db *mgo.Database) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "PUT" {
			r.ParseForm()
			values := r.Form
			id := values.Get("id")
			title := values.Get("title")
			content := values.Get("content")
			c := db.C("feeds")
			var feed entity.Feed
			if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&feed); err == nil {
				if title != "" {
					feed.Title = title
				}
				if content != "" {
					feed.Content = content
				}
				c.UpdateId(bson.ObjectIdHex(id), feed)
				w.WriteHeader(200)
				return
			}
			w.WriteHeader(500)
		} else if r.Method == "OPTIONS" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(401)
		}
	}
}