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
	begin, err2 := p.db.Begin()
	if err2 != nil {
		return err2
	}

	stmt, err := begin.Prepare("insert into public.\"group\"(uuid, title, description) values ($1,$2,$3)")
	if err != nil {
		log.Printf("Err build Prepare: %v", err)
		return err
	}
	_, err = stmt.Exec(group.ID().String(), group.Title(), group.Description())
	if err != nil {
		log.Printf("Err Exec: %v", err)
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return err
	}

	tasksErr := p.saveTasksIdsByGroupId(begin, group.ID().String(), group.Tasks())
	if tasksErr != nil {
		log.Printf("Err Exec: %v", tasksErr)
		err := begin.Rollback()
		if err != nil {
			return err
		}
		return tasksErr
	}

	err = begin.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresGroup) Update(group *domain.Group) error {
	panic("implement me")
}

func (p *PostgresGroup) GetByID(id domain.GroupID) (*domain.Group, error) {
	query := fmt.Sprintf(
		"select uuid, title, description from public.\"group\" where uuid like '%v'", id.String())

	rows, err := p.db.Query(query)
	if err != nil {
		log.Printf("Error Query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var uuidStr string
		var title string
		var description string
		var taskIDs []domain.TaskID

		err := rows.Scan(&uuidStr, &title, &description)
		if err != nil {
			log.Printf("Error Rows Scan: %v", err)
			return nil, err
		}

		taskIDs, taskErr := p.findAllTaskIDsByGroupId(uuidStr)
		if taskErr != nil {
			log.Printf("Find TaskIds by Group id: %v", taskErr)
			return nil, taskErr
		}

		uuid, err := domain.NewGroupIDFromString(uuidStr)

		if err != nil {
			log.Printf("Error parsing UUID %v. Err: %v", uuidStr, err)
		} else {
			return domain.BuildGroup(uuid, title, description, taskIDs), nil
		}
	}

	return nil, nil
}

func (p *PostgresGroup) DelByID(id domain.GroupID) error {
	panic("implement me")
}

func (p *PostgresGroup) GetAll() ([]*domain.Group, error) {
	groups := make([]*domain.Group, 0, 10)

	rows, err := p.db.Query("select uuid, title, description from public.\"group\"")
	if err != nil {
		log.Printf("Error Query: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var uuidStr string
		var title string
		var description string
		var taskIDs []domain.TaskID

		err := rows.Scan(&uuidStr, &title, &description)
		if err != nil {
			log.Printf("Error Rows Scan: %v", err)
			return nil, err
		}

		taskIDs, taskErr := p.findAllTaskIDsByGroupId(uuidStr)
		if taskErr != nil {
			log.Printf("Find TaskIds by Group id: %v", taskErr)
			return nil, taskErr
		}

		uuid, err := domain.NewGroupIDFromString(uuidStr)

		if err != nil {
			log.Printf("Error parsing UUID %v. Err: %v", uuidStr, err)
		} else {
			group := domain.BuildGroup(uuid, title, description, taskIDs)
			groups = append(groups, group)
		}
	}

	return groups, nil
}

func (p *PostgresGroup) saveTasksIdsByGroupId(begin *sql.Tx, groupId string, taskIds []domain.TaskID) error {
	stmt, err := begin.Prepare("insert into public.group_task(group_uuid, task_uuid) VALUES ($1, $2)")
	if err != nil {
		log.Printf("Err build Prepare: %v", err)
		return err
	}

	for _, taskId := range taskIds {
		_, err = stmt.Exec(groupId, taskId.String())
		if err != nil {
			log.Printf("Err Exec: %v", err)
			return err
		}
	}

	return nil
}

func (p *PostgresGroup) findAllTaskIDsByGroupId(groupId string) ([]domain.TaskID, error) {
	taskIds := make([]domain.TaskID, 0, 10)

	query := fmt.Sprintf("select task_uuid from public.group_task where group_uuid like '%v'", groupId)
	rows, err := p.db.Query(query)
	if err != nil {
		log.Printf("Error Query: %v", err)
		return nil, err
	}
	defer rows.Close()

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
