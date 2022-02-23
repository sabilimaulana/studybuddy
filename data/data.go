package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error
	db, err = sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		return err
	}
	return db.Ping()
}

func CreateTable() {
	CreateTableSQL := `CREATE TABLE IF NOT EXISTS studybuddy (
		"idNote" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	statement, err := db.Prepare(CreateTableSQL)

	if err != nil {
		panic(err)
	}
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	log.Println("Studybuddy table created")
}

func InsertNote(word string, definition string, category string) {
	insertNoteSQL := `INSERT INTO studybuddy (word, definition, category) VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertNoteSQL)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(word, definition, category)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Note inserted")
}

func DisplayAllNotes() {
	row, err := db.Query("SELECT * FROM studybuddy ORDER BY idNote")
	if err != nil {
		log.Fatalln(err)
	}

	defer row.Close()

	for row.Next() {
		var idNote int
		var word, definition, category string

		err = row.Scan(&idNote, &word, &definition, &category)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("%d. [%s] %s, %s\n", idNote, category, word, definition)
	}

}
