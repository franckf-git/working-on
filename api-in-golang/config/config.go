package config

var PORT string = ":8000"
var WelcomeMessage string = "Welcome, to have more details about this API, visit /api/v1/docs"
var Database string = "./storage/database.sqlite3"
var DocsLink string = "https://gitlab.com/franckf/working-on/-/blob/master/api-in-golang/readme.md#documentation-de-lapi"
var JWTkey string = "2d01d5d9c24034d54fe4fba0ede5182d"

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Datas   string `json:"datas"`
	Created string `json:"created"`
	IdUser  int    `json:"idUser"`
}

type NewPost struct {
	Title  string `json:"title"`
	Datas  string `json:"datas"`
	IdUser int    `json:"idUser"`
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Id      int    `json:"id"`
}
