package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

var db *sql.DB

func TestNewPostgresGroup(t *testing.T) {
	_ = NewPostgresGroup(db)
}

func init() {
	user := "postgres"
	password := "postgres"
	dbname := "stepanlahov"
	dbtype := "postgres"

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		user, password, dbname)

	var err error

	db, err = sql.Open(dbtype, connStr)
	if err != nil {
		panic(err)
	}

	log.Printf("Connection!!!!!")
}
