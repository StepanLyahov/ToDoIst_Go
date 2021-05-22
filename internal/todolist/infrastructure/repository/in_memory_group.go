package repository

import (
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	"github.com/pkg/errors"
)

type InMemoryGroup struct {
	db map[domain.GroupID]*domain.Group
}

func NewInMemoryGroup() InMemoryGroup {
	db := make(map[domain.GroupID]*domain.Group)

	return InMemoryGroup{
		db: db,
	}
}

func (i InMemoryGroup) Save(group *domain.Group) error {
	i.db[group.ID()] = group
	return nil
}

func (i InMemoryGroup) Update(group *domain.Group) error {
	_, err := i.db[group.ID()]
	if err == false {
		return errors.New("not found")
	}

	i.db[group.ID()] = group
	return nil
}

func (i InMemoryGroup) GetByID(id domain.GroupID) (*domain.Group, error) {
	_, err := i.db[id]
	if err == false {
		return &domain.Group{}, errors.New("not found")
	}

	return i.db[id], nil
}

func (i InMemoryGroup) DelByID(id domain.GroupID) error {
	_, err := i.db[id]
	if err == false {
		return errors.New("not found")
	}

	delete(i.db, id)
	return nil
}

func (i InMemoryGroup) GetAll() []*domain.Group {
	var groups []*domain.Group

	for _, group := range i.db {
		groups = append(groups, group)
	}
	return groups
}
