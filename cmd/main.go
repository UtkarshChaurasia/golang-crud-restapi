package main

import (
	"database/sql"
	"golang-restapi/platform/newsfeed"

	_ "github.com/mattn/go-sqlite3"
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

	feed.Add(newsfeed.Item{
		Content: "Hello, Utkarsh!",
	})

}
