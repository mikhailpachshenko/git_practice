package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mikhailpachshenko"
	password = "1722"
	dbname   = "project"
)

func testField(field []int) {
	for i := range field {
		field[i] = i
	}
}

func ConnectionDatabase() *sql.DB {
	pgSqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pgSqlConn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
