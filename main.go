package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	const sqliteFile = "cass_notes.sqlite"

	db, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		fmt.Println("Error opening sqlite db", err.Error())
	}
	defer func() {
	    er := db.Close()
		if er != nil {
			fmt.Println("Error closing DB", er.Error())
		}
	}()

	type Note struct {
		Id int64 `json:"id"`
		Title string `json:"title"`
		Description string `json:"description"`
	}
}
