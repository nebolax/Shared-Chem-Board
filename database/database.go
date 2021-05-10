package database

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	mu = sync.Mutex{}
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "ShChemBoard"
)

type DBI struct {
	ID    interface{}
	Login interface{}
	Email interface{}
}

func loadVal(rows *sql.Rows, exval interface{}) []interface{} {
	rf := reflect.ValueOf(exval)
	res := []interface{}{}

	for rows.Next() {
		ar := make([]interface{}, rf.NumField())
		ar1 := make([]interface{}, rf.NumField())
		for i := 0; i < rf.NumField(); i++ {
			ar1[i] = &ar[i]
		}
		if err := rows.Scan(ar1...); err != nil {
			panic(err)
		}

		mr := reflect.New(reflect.Indirect(reflect.ValueOf(exval)).Type()).Elem()

		for i := 0; i < rf.NumField(); i++ {
			mr.Field(i).Set(reflect.ValueOf(ar[i]))
		}
		res = append(res, mr.Interface())

	}

	return res
}

func Query(query string, exval interface{}) []interface{} {
	mu.Lock()
	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	mu.Unlock()
	return loadVal(rows, exval)
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
}
