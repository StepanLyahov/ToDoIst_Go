package command

import (
	"app/app/query"
	"app/app/repository"
	"app/domain"
)

type ChangeTaskHandler struct {
	taskRepos repository.TaskRepository
}

func NewChangeTaskHandler(taskRepos repository.TaskRepository) ChangeTaskHandler {
	return ChangeTaskHandler{
		taskRepos: taskRepos,
	}
}

func (h ChangeTaskHandler) Execute(taskDto query.TaskDto) error {
	taskID, err := domain.NewTaskIDFromString(taskDto.Id)
	if err != nil {
		return err
	}

	_, err = h.taskRepos.GetByID(taskID)
	if err != nil {
		return err
	}

	updatedTask, err := query.TaskToDomain(taskDto)
	if err != nil {
		return err
	}

	err = h.taskRepos.Update(&updatedTask)
	if err != nil {
		return err
	}
	return nil
}
