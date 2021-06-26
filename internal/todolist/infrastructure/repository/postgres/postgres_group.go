package postgres

import (
	"database/sql"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	_ "github.com/lib/pq"
	"log"
)

type PostgresGroup struct {
	db *sql.DB
}

func NewPostgresGroup(db *sql.DB) *PostgresGroup {
	return &PostgresGroup{db: db}
}

func (p *PostgresGroup) Save(group *domain.Group) error {
	panic("implement me")
}

func (p *PostgresGroup) Update(group *domain.Group) error {
	panic("implement me")
}

func (p *PostgresGroup) GetByID(id domain.GroupID) (*domain.Group, error) {
	panic("implement me")
}

func (p *PostgresGroup) DelByID(id domain.GroupID) error {
	panic("implement me")
}

func (p *PostgresGroup) GetAll() []*domain.Group {
	groups := make([]*domain.Group, 0, 10)

	rows, err := p.db.Query("select uuid, title, description from public.\"group\"")
	if err != nil {
		log.Printf("Error Query: %v", err)
		return nil
	}

	for rows.Next() {
		var uuidStr string
		var title string
		var description string
		var taskIDs []domain.TaskID

		err := rows.Scan(&uuidStr, &title, &description)
		if err != nil {
			log.Printf("Error Rows Scan: %v", err)
			return nil
		}

		taskIDs, err = p.findAllTaskIDsByGroupId(uuidStr)

		uuid, err := domain.NewGroupIDFromString(uuidStr)

		if err != nil {
			log.Printf("Error parsing UUID %v. Err: %v", uuidStr, err)
		} else {
			group := domain.Build(uuid, title, description, taskIDs)
			groups = append(groups, group)
		}
	}

	return groups
}

func (p *PostgresGroup) findAllTaskIDsByGroupId(groupId string) ([]domain.TaskID, error) {
	taskIds := make([]domain.TaskID, 0, 10)

	query := fmt.Sprintf("select task_uuid from public.group_task where group_uuid like '%v'", groupId)
	rows, err := p.db.Query(query)
	if err != nil {
		log.Printf("Error Query: %v", err)
		return nil, err
	}

	for rows.Next() {
		var uuidStr string

		err := rows.Scan(&uuidStr)
		if err != nil {
			log.Printf("Error Rows Scan: %v", err)
			return nil, err
		}

		uuid, err := domain.NewTaskIDFromString(uuidStr)

		if err != nil {
			log.Printf("Error parsing UUID %v. Err: %v", uuidStr, err)
		} else {
			taskIds = append(taskIds, uuid)
		}
	}

	return taskIds, nil
}
