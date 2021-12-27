package newsfeed

import "database/sql"

type Feed struct {
	DB *sql.DB // pointer to a sql db
}

// func (feed *Feed) Get() []Item {

// }

func (feed *Feed) Add(item Item) {
	stmt, _ := feed.DB.Prepare(`
		INSERT INTO newsfeed (content) values (?)
	`)
	stmt.Exec(item.Content)

}

// constructor function for feed returns pointer to feed
func NewFeed(db *sql.DB) *Feed {

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

	return &Feed{
		DB: db,
	}
}
