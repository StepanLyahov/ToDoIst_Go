package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
)

func TestNewPostgresGroup(t *testing.T) {
	db, err := initPostgresConnection()
	if err != nil {
		panic(err)
	}

	_ = NewPostgresGroup(db)
}

func TestPostgresGroup_GetAll(t *testing.T) {
	db, err := initPostgresConnection()
	if err != nil {
		panic(err)
	}

	rep := NewPostgresGroup(db)

	res, err := rep.GetAll()
	if err != nil {
		t.Fatalf("Err must be nil, but %v", err)
	}

	for _, g := range res {
		log.Printf("Group['%v' '%v' '%v', Tasks {%v}]", g.ID(), g.Title(), g.Description(), g.Tasks())
	}
}

func initPostgresConnection() (*sql.DB, error) {
	user := "postgres"
	password := "postgres"
	dbname := "stepanlahov"
	dbtype := "postgres"

	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
		user, password, dbname)

	db, err := sql.Open(dbtype, connStr)

	log.Print("Connected!!!")

	return db, err
}
