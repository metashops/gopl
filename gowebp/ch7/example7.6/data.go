package example7_6

import (
	`database/sql`
	`time`

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

func NewPost() *Post {
	return &Post{}
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func init() {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/db-hello-02")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

func retrieve(id int) (post *Post, err error) {
	post = NewPost()
	DB.QueryRow("select id, content, author from posts where id = $1", id).
		Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (p *Post) create() (err error) {
	statement := "insert into posts (content,author) values ($1, $2) returning id"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(p.Content, p.Author).Scan(&p.Id)
	return
}

func (p *Post) update() (err error) {
	_, err = DB.Exec("update posts set content = $2, author = $3 where id = $1", p.Id, p.Content, p.Author)
	return
}

func (p *Post) delete() (err error) {
	_, err = DB.Exec("delete from posts where id = $1", p.Id)
	return
}
