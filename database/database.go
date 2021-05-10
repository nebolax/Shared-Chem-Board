package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "TestDB" //"ShChemBoard"
)

type DBI struct {
	ID    interface{}
	Login interface{}
	Email interface{}
}

func observerWorker(inp chan string, out chan *sql.Rows) {
	query := <-inp
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	out <- rows
}

func NewDbObserver() (chan string, chan *sql.Rows) {
	inp := make(chan string)
	out := make(chan *sql.Rows)
	go observerWorker(inp, out)
	return inp, out
}

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// rows, err := db.Query(``)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// users := make([]DBI, 0)

	// for rows.Next() {
	// 	var id interface{}
	// 	var login interface{}
	// 	var email interface{}
	// 	if err := rows.Scan(&id, &login, &email); err != nil {
	// 		panic(err)
	// 	}
	// 	user := DBI{id, login, email}
	// 	users = append(users, user)
	// }
	// fmt.Printf("%v", users)

	// fmt.Println("Successfully connected to db")
}
