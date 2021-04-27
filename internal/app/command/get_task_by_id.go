package command

import (
	"app/repository"
	"domain"
)

type GetTaskByIdHandler struct {
	taskRepos repository.TaskRepository
}

func NewGetTaskByIdHandler(rep repository.TaskRepository) GetTaskByIdHandler {
	return GetTaskByIdHandler{taskRepos: rep}
}

func (h GetTaskByIdHandler) Execute(taskUuid string) (*domain.Task, error) {
	taskID, err := domain.NewTaskIDFromString(taskUuid)
	if err != nil {
		return &domain.Task{}, err
	}

	task, err := h.taskRepos.GetByID(taskID)
	if err != nil {
		return &domain.Task{}, err
	}

	return task, nil
}