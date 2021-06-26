package postgres

import (
	"database/sql"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	_ "github.com/lib/pq"
)

type PostgresGroup struct {
	db *sql.DB
}

func NewPostgresGroup(db *sql.DB) *PostgresGroup {
	return &PostgresGroup{db: db}
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
