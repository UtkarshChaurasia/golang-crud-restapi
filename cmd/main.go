package main

import (
	"database/sql"
	"golang-restapi/platform/newsfeed"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qkgo/yin"
)

/*
CREATE TABLE IF NOT EXISTS "newsfeed" (
	"ID"	INTEGER NOT NULL,
	"content"	TEXT,
	PRIMARY KEY("ID" AUTOINCREMENT)
);
*/

func main() {

	// creates a database of specified driver and filename
	db, _ := sql.Open("sqlite3", "./newsfeed.db")
	feed := newsfeed.NewFeed(db)

	// create a router
	r := chi.NewRouter()

	// ROUTE 1: Get all the posts using: GET "/posts"
	r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
		res, _ := yin.Event(w, r)
		items := feed.Get()
		res.SendJSON(items)
	})

	// ROUTE 2: ADD the post using: POST "/addpost"
	r.Post("/addpost", func(w http.ResponseWriter, r *http.Request) {

		res, req := yin.Event(w, r)
		body := map[string]string{}
		req.BindBody(&body)
		item := newsfeed.Item{
			Content: body["content"],
		}
		feed.Add(item)
		res.SendJSON(item)

	})

	http.ListenAndServe(":3000", r)

}
