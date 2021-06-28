package postgres

import (
	"database/sql"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
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

func TestPostgresGroup_Save(t *testing.T) {
	db, err := initPostgresConnection()
	if err != nil {
		panic(err)
	}

	rep := NewPostgresGroup(db)

	test := domain.NewGroup("new", "new")
	err = test.AddTask(domain.NewTaskID())
	if err != nil {
		t.Fatalf("Err must be nil, but %v", err)
	}

	err = rep.Save(test)
	if err != nil {
		t.Fatalf("Err must be nil, but %v", err)
	}
}

func TestPostgresGroup_GetByID(t *testing.T) {
	db, err := initPostgresConnection()
	if err != nil {
		panic(err)
	}

	rep := NewPostgresGroup(db)

	fromString, err := domain.NewGroupIDFromString("f8c67835-49ad-4f3e-9e49-4fdb3d62b4e5")
	if err != nil {
		t.Fatalf("Err must be nil, but err: %v", err)
	}

	g, err := rep.GetByID(fromString)
	if err != nil {
		t.Fatalf("Err must be nil, but err: %v", err)
	}

	if g == nil {
		t.Fatal("Group must be exist")
	}

	log.Printf("Group['%v' '%v' '%v', Tasks {%v}]", g.ID(), g.Title(), g.Description(), g.Tasks())
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
