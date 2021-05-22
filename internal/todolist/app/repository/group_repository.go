package repository

import "todolist/domain"

type GroupRepository interface {
	Save(group *domain.Group) error
	Update(group *domain.Group) error
	GetByID(id domain.GroupID) (*domain.Group, error)
	DelByID(id domain.GroupID) error
	GetAll() []*domain.Group
}
