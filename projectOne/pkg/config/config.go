package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// same
const (
	host     = "localhost"
	port     = 5432
	user     = "mikhailpachshenko"
	password = "1722"
)

func fieldNotNull(input []int) bool {
	if len(input) != 0 {
		return true
	}
	return false
}

func fieldNotNull(input []int) bool {
	if len(input) != 0 {
		return true
	}
	return false
}

func fieldNotNull(input []int) bool {
	if len(input) != 0 {
		return true
	}
	return false
}

func testField(field []int) {
	for i := range field {
		field[i] = i
	}
}

func ConnectionDatabase() *sql.DB {
	pgSqlConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=project sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", pgSqlConn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
