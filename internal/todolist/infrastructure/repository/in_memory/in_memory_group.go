package in_memory

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
	_, ok := i.db[group.ID()]
	if !ok {
		return errors.New("not found")
	}

	i.db[group.ID()] = group
	return nil
}

func (i InMemoryGroup) GetByID(id domain.GroupID) (*domain.Group, error) {
	_, ok := i.db[id]
	if !ok {
		return &domain.Group{}, errors.New("not found")
	}

	return i.db[id], nil
}

func (i InMemoryGroup) DelByID(id domain.GroupID) error {
	_, ok := i.db[id]
	if !ok {
		return errors.New("not found")
	}

	delete(i.db, id)
	return nil
}

func (i InMemoryGroup) GetAll() ([]*domain.Group, error) {
	var groups []*domain.Group

	for _, group := range i.db {
		groups = append(groups, group)
	}
	return groups, nil
}
