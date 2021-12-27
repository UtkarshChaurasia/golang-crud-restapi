package newsfeed

import "database/sql"

type Feed struct {
	DB *sql.DB // pointer to a sql db
}

func (feed *Feed) Get() []Item {

	// query statement stores fetched result in rows
	rows, _ := feed.DB.Query(`
		SELECT * FROM newsfeed
	`)

	// new slice of Item type to store the rows
	items := []Item{}

	var id int
	var content string

	// iterating through all rows
	for rows.Next() {

		// get id and content and store it in var id & var content
		rows.Scan(&id, &content)

		// create a new Item with id and content
		newItem := Item{
			ID:      id,
			Content: content,
		}

		// append newItem to items slice
		items = append(items, newItem)
	}

	return items

}

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
