package handler

import (
	"github.com/gaoyuyue/feed-backend/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAddFeed(t *testing.T)  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
	db := session.DB("feed")

	tests := []struct {
		method string
		body   string
		code   int
	}{
		{"POST", "title=test&content=helloworld", 200},
		{"GET", "title=test&content=helloworld", 401},
		{"DELETE", "title=test&content=helloworld", 401},
		{"PUT", "title=test&content=helloworld", 401},
	}
	for _, test := range tests {
		request, err := http.NewRequest(test.method, "/add_one", strings.NewReader(test.body))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		AddFeed(db)(recorder, request)

		if recorder.Code != test.code {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Code, test.code)
		}
	}
}

func TestDeleteFeed(t *testing.T)  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
	db := session.DB("feed")

	id := bson.NewObjectId()
	db.C("feeds").Insert(&entity.Feed{Id: id,
		Title:"test", Content:"test", CreateTime:time.Now()})

	tests := []struct {
		method string
		url string
		code   int
	}{
		{"DELETE", "/delete_one?id="+id.Hex(), 200},
		{"GET", "/delete_one?id="+id.Hex(), 401},
		{"POST", "/delete_one?id="+id.Hex(), 401},
		{"PUT", "/delete_one?id="+id.Hex(), 401},
	}
	for _, test := range tests {
		request, err := http.NewRequest(test.method, test.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		DeleteFeed(db)(recorder, request)

		if recorder.Code != test.code {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Code, test.code)
		}
	}
}

func TestUpdateFeed(t *testing.T)  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
	db := session.DB("feed")

	id := bson.NewObjectId()
	db.C("feeds").Insert(&entity.Feed{Id: id,
		Title:"test", Content:"test", CreateTime:time.Now()})

	tests := []struct {
		method string
		body   string
		code   int
	}{
		{"PUT", "id=" + id.Hex() + "&title=updatetest&content=updatetest", 200},
		{"GET", "id=" + id.Hex() + "&title=updatetest&content=updatetest", 401},
		{"DELETE", "id=" + id.Hex() + "&title=updatetest&content=updatetest", 401},
		{"POST", "id=" + id.Hex() + "&title=updatetest&content=updatetest", 401},
	}
	for _, test := range tests {
		request, err := http.NewRequest(test.method, "/update_one", strings.NewReader(test.body))
		request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		UpdateFeed(db)(recorder, request)

		if recorder.Code != test.code {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Code, test.code)
		}
	}
}

func TestListFeed(t *testing.T)  {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		t.Error(err)
	}
	defer session.Close()
	db := session.DB("feed")

	tests := []struct {
		method string
		url string
		code   int
	}{
		{"GET", "/list", 200},
		{"PUT", "/list", 401},
		{"DELETE", "/list", 401},
		{"POST", "/list", 401},
	}
	for _, test := range tests {
		request, err := http.NewRequest(test.method, test.url, nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		ListFeed(db)(recorder, request)

		if recorder.Code != test.code {
			t.Errorf("handler returned unexpected body: got %v want %v",
				recorder.Code, test.code)
		}
	}
}