package database

import (
	"database/sql"
	"fmt"
	"math"
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

func Uint8To64Ar(inp []uint8) (res []uint64) {

	res = []uint64{}
	for i := 0; i < len(inp); i++ {
		if i%8 == 0 {
			res = append(res, 0)
		}
		res[i/8] += uint64(inp[i]) * uint64(math.Pow(256, 7-float64(i%8)))
		fmt.Printf("%v\n", res)
	}

	return
}

func Uint64To8Ar(inp []uint64) (res []uint8) {
	res = []uint8{}
	for i := 0; i < len(inp); i++ {
		buf := make([]uint8, 8)
		for a := 0; a < 8; a++ {
			buf[7-a] = uint8(inp[i] % 256)
			inp[i] = inp[i] / 256
		}
		res = append(res, buf...)
	}

	return
}

func loadVal(rows *sql.Rows, exval interface{}) []interface{} {
	rf := reflect.ValueOf(exval)
	res := []interface{}{}

	for rows.Next() {
		mr := reflect.New(reflect.Indirect(reflect.ValueOf(exval)).Type()).Elem()

		ar := make([]interface{}, rf.NumField())

		ar1 := make([]interface{}, rf.NumField())
		for i := 0; i < rf.NumField(); i++ {
			ar1[i] = &ar[i]
		}
		if err := rows.Scan(ar1...); err != nil {
			panic(err)
		}

		for i := 0; i < rf.NumField(); i++ {
			switch reflect.TypeOf(ar[i]) {
			case reflect.TypeOf([]uint8{}):
				// v := ar[i].([]uint8)
				// mr.Field(i).Set(reflect.ValueOf((pq.Int64Array)(v)))
			default:
				mr.Field(i).Set(reflect.ValueOf(ar[i]))
			}
		}
		res = append(res, mr.Interface())

	}

	return res
}

func Query(query string, exval interface{}) []interface{} {
	// mu.Lock()
	// rows, err := db.Query(query)
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// mu.Unlock()
	// return loadVal(rows, exval)
	return nil
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
