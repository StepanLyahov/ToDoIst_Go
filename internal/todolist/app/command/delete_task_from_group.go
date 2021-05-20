package command

import (
	"app/app/repository"
	"app/domain"
)

type DeleteTaskFromGroupHandler struct {
	groupRepos repository.GroupRepository
	taskRepos  repository.TaskRepository
}

func NewDeleteTaskFromGroupHandler(groupRepos repository.GroupRepository,
	taskRepos repository.TaskRepository) DeleteTaskFromGroupHandler {
	return DeleteTaskFromGroupHandler{
		groupRepos: groupRepos,
		taskRepos:  taskRepos,
	}
}

func (h DeleteTaskFromGroupHandler) Execute(groupUuid string, taskUuid string) error {
	group, err := h.getGroupByID(groupUuid)
	if err != nil {
		return err
	}

	task, errDb := h.getTaskByID(taskUuid)
	if errDb != nil {
		return errDb
	}

	group.DelTask(task.ID())

	errRepos := h.groupRepos.Save(group)
	if errRepos != nil {
		return errRepos
	}

	return nil
}

func (h DeleteTaskFromGroupHandler) getGroupByID(groupUuid string) (*domain.Group, error) {
	groupID, err := domain.NewGroupIDFromString(groupUuid)
	if err != nil {
		return &domain.Group{}, err
	}

	group, err := h.groupRepos.GetByID(groupID)
	if err != nil {
		return &domain.Group{}, err
	}

	return group, nil
}

func (h DeleteTaskFromGroupHandler) getTaskByID(taskUuid string) (*domain.Task, error) {
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
