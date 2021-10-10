package models

import (
	"database/sql"
	"errors"
	"time"

	"lite-api-crud/config"

	_ "github.com/mattn/go-sqlite3"
)

func RegisterPost(db *sql.DB, title string, datas string, idUser int) (id int, err error) {
	stmt, err := db.Prepare("INSERT INTO posts(title, datas, created, idUser) VALUES(?, ?, ?, ?)")
	if err != nil {
		config.ErrorLogg("RegisterPost(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	var created string = time.Now().Format(time.RFC3339)
	result, err := stmt.Exec(title, datas, created, idUser)
	if err != nil {
		config.ErrorLogg("RegisterPost(models) - executing query:", err)
		return
	}

	idReturn, _ := result.LastInsertId()
	id = int(idReturn)
	return
}

func GetAllPosts(db *sql.DB) (Posts []config.GetPost, err error) {
	stmt, err := db.Prepare("SELECT * FROM posts")
	if err != nil {
		config.ErrorLogg("GetAllPosts(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		config.ErrorLogg("GetAllPosts(models) - executing query:", err)
		return
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
			config.ErrorLogg("GetAllPosts(models) - scanning values:", err)
			return
		}
		post := config.Post{
			Title:  title,
			Datas:  datas,
			IdUser: id,
		}
		currentPost := config.GetPost{
			Id:      id,
			Post:    post,
			Created: created,
		}
		Posts = append(Posts, currentPost)
	}
	err = rows.Err()
	if err != nil {
		config.ErrorLogg("GetAllPosts(models) - reading rows:", err)
		return
	}
	return
}

func GetPost(db *sql.DB, id int) (Post config.GetPost, err error) {
	stmt, err := db.Prepare("SELECT * FROM posts WHERE id = ?")
	if err != nil {
		config.ErrorLogg("GetPost(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	var title string
	var datas string
	var created string
	var idUser int
	err = stmt.QueryRow(id).Scan(&id, &title, &datas, &created, &idUser)
	if err != nil {
		config.ErrorLogg("GetPost(models) - reading rows:", err)
		return
	}
	post := config.Post{
		Title:  title,
		Datas:  datas,
		IdUser: id,
	}
	Post = config.GetPost{
		Id:      id,
		Post:    post,
		Created: created,
	}
	return
}

func UpdatingPost(db *sql.DB, id int, title string, datas string, idUser int) (err error) {
	stmt, err := db.Prepare("UPDATE posts SET title = ?, datas = ? WHERE idUser = ? AND id=?")
	if err != nil {
		config.ErrorLogg("UpdatingPost(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(title, datas, idUser, id)
	if err != nil {
		config.ErrorLogg("UpdatingPost(models) - executing query:", err)
		return
	}

	lines, _ := result.RowsAffected()
	if lines == 0 {
		config.ErrorLogg("UpdatingPost(models) - id not found")
		return errors.New("update fail - id not found")
	}
	return
}

func DeletingPost(db *sql.DB, id int) (err error) {
	stmt, err := db.Prepare("DELETE FROM posts WHERE id = ?")
	if err != nil {
		config.ErrorLogg("DeletingPost(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		config.ErrorLogg("DeletingPost(models) - executing query:", err)
		return
	}

	lines, _ := result.RowsAffected()
	if lines == 0 {
		config.ErrorLogg("DeletingPost(models) - id not found")
		return errors.New("delete fail - id not found")
	}
	return
}

func GetAllPostsByUser(db *sql.DB, idUser int) (ids []int, err error) {
	stmt, err := db.Prepare("SELECT id FROM posts WHERE idUser = ?")
	if err != nil {
		config.ErrorLogg("GetAllPosts(models) - preparing query:", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(idUser)
	if err != nil {
		config.ErrorLogg("GetAllPostsByUser(models) - executing query:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			config.ErrorLogg("GetAllPostsByUser(models) - scanning values:", err)
			return
		}
		ids = append(ids, id)
	}
	err = rows.Err()
	if err != nil {
		config.ErrorLogg("GetAllPostsByUser(models) - reading rows:", err)
		return
	}
	return
}
