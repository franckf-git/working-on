package models

import (
	"database/sql"
	"log"
	"time"

	"lite-api-crud/config"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterPost(db *sql.DB, title string, datas string, idUser int) (id int, err error) {
	var created string = time.Now().Format(time.RFC3339)

	insert, err := db.Begin()
	if err != nil {
		log.Fatal("Insert fail - opening database:", err)
		return 0, err
	}
	stmt, err := insert.Prepare("INSERT INTO posts(title, datas, created, idUser) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Insert fail - preparing query:", err)
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(title, datas, created, idUser)
	if err != nil {
		log.Println("Insert fail - executing query:", err)
		return 0, err
	}
	insert.Commit()

	idReturn, _ := result.LastInsertId()
	id = int(idReturn)
	return
}

func GetAllPosts(db *sql.DB) (Posts []config.Post) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal("Select fail - executing query:", err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var title string
		var datas string
		var created string
		var idUser int
		err = rows.Scan(&id, &title, &datas, &created, &idUser)
		if err != nil {
			log.Fatal("Select fail - scanning values:", err)
		}
		currentPost := config.Post{
			Id:      id,
			Title:   title,
			Datas:   datas,
			Created: created,
			IdUser:  idUser,
		}
		Posts = append(Posts, currentPost)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Select fail - reading rows:", err)
	}
	return
}
