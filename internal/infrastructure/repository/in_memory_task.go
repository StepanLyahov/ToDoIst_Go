package repository

import (
	"domain"
	"errors"
)

type InMemoryTask struct {
	db map[domain.TaskID]*domain.Task
}

func (i InMemoryTask) Save(task *domain.Task) error {
	i.db[task.ID()] = task
	return nil
}

func (i InMemoryTask) GetByID(id domain.TaskID) (*domain.Task, error) {
	task, err := i.db[id]
	if err == false {
		return &domain.Task{}, errors.New("not found")
	}

	return task, nil
}

func (i InMemoryTask) DelByID(id domain.TaskID) error {
	_, err := i.db[id]
	if err == false {
		return errors.New("not found")
	}

	delete(i.db, id)
	return nil
}
