package entity

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Feed struct {
	Id bson.ObjectId `json:"id" bson:"_id"`
	Title string `json:"title"`
	Content string `json:"content"`
	//Author string `json:"author"`
	CreateTime time.Time `json:"create_time"`
}
