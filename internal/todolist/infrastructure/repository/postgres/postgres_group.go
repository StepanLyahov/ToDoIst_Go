package postgres

import (
	"database/sql"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

var (
	once sync.Once

	instance *PostgresGroup = nil
)

type PostgresGroup struct {
	user     string
	password string
	dbname   string
	dbtype   string
	db       sql.DB
}

func NewPostgresGroup(user, password, dbname string) *PostgresGroup {
	once.Do(func() {
		rep := PostgresGroup{
			user:     user,
			password: password,
			dbname:   dbname,
			dbtype:   "postgres",
		}

		connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
			rep.user, rep.password, rep.dbname)

		db, err := sql.Open(rep.dbtype, connStr)
		if err != nil {
			panic(err)
		}
		log.Print("Connected!!!")
		defer db.Close()

		instance = &rep
	})

	return instance
}

func (p PostgresGroup) Save(group *domain.Group) error {
	panic("implement me")
}

func (p PostgresGroup) Update(group *domain.Group) error {
	panic("implement me")
}

func (p PostgresGroup) GetByID(id domain.GroupID) (*domain.Group, error) {
	panic("implement me")
}

func (p PostgresGroup) DelByID(id domain.GroupID) error {
	panic("implement me")
}

func (p PostgresGroup) GetAll() []*domain.Group {
	panic("implement me")
}
