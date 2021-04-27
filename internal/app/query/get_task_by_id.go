package query

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

func (h GetTaskByIdHandler) Execute(taskUuid string) (TaskDto, error) {
	taskID, err := domain.NewTaskIDFromString(taskUuid)
	if err != nil {
		return TaskDto{}, err
	}

	task, err := h.taskRepos.GetByID(taskID)
	if err != nil {
		return TaskDto{}, err
	}

	taskDto := TaskToDto(*task)

	return taskDto, nil
}