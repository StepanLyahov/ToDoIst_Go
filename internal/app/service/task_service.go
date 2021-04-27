package service

import (
	"app/repository"
	"domain"
)

type TaskService struct {
	taskRepos  repository.TaskRepository
}

func NewTaskService(taskRepos repository.TaskRepository) *TaskService {
	return &TaskService{taskRepos: taskRepos}
}

func (ts *TaskService) Create(title string, description string, priority uint8) (domain.TaskID, error) {

	pr, err := domain.NewPriorityFromUint8(priority)
	if err != nil {
		return domain.TaskID{}, err
	}

	task := domain.NewTaskWithCurrentDate(title, description, pr)

	errTask := ts.taskRepos.Save(task)
	if errTask != nil {
		return domain.TaskID{}, errTask
	}

	return task.ID(), nil
}
