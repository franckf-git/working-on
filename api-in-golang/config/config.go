package config

var PORT string = ":8000"
var WelcomeMessage string = "Welcome, to have more details about this API, visit /api/v1/docs"
var Database string = "./storage/database.sqlite3"

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Datas   string `json:"datas"`
	Created string `json:"created"`
	IdUser  int    `json:"idUser"`
}
