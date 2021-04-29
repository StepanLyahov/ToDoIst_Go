package repository

import (
	"domain"
	"errors"
)

type InMemoryTask struct {
	db map[domain.TaskID]*domain.Task
}

func NewInMemoryTask() InMemoryTask {
	db := make(map[domain.TaskID]*domain.Task)

	return InMemoryTask {
		db: db,
	}
}

func (i InMemoryTask) Save(task *domain.Task) error {
	i.db[task.ID()] = task
	return nil
}

func (i InMemoryTask) Update(task *domain.Task) error {
	_, err := i.db[task.ID()]
	if err == false {
		return errors.New("not found")
	}

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
