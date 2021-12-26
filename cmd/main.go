package main

import (
	"database/sql"

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

	// prepares sql statement
	stmt, _ := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "newsfeed" (
		"ID"	INTEGER NOT NULL,
		"content"	TEXT,
		PRIMARY KEY("ID" AUTOINCREMENT)
	);
	`)

	// executes sql statement
	stmt.Exec()
}
