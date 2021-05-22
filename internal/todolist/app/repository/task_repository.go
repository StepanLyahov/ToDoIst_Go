package repository

import "github.com/StepanLyahov/ToDoIst/todolist/domain"

type TaskRepository interface {
	Save(task *domain.Task) error
	Update(task *domain.Task) error
	GetByID(id domain.TaskID) (*domain.Task, error)
	DelByID(id domain.TaskID) error
}
