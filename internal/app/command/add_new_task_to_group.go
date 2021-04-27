package command

import (
	"app/query"
	"app/repository"
	"domain"
)

type AddNewTaskToGroupHandler struct {
	groupRepos repository.GroupRepository
	taskRepos  repository.TaskRepository
}


func NewAddNewTaskToGroupHandler(groupRepos repository.GroupRepository,
	taskRepos repository.TaskRepository) AddNewTaskToGroupHandler {
	return AddNewTaskToGroupHandler{
		groupRepos: groupRepos,
		taskRepos:  taskRepos,
	}
}

func (h AddNewTaskToGroupHandler) Execute(groupUuid string, taskDto query.TaskDto) (domain.TaskID, error) {
	taskId, err := h.createTaskAndSave(taskDto)
	if err != nil {
		return domain.TaskID{}, err
	}

	errTask := h.addNewTaskToGroup(groupUuid, taskId)
	if errTask != nil {
		return domain.TaskID{}, errTask
	}

	return taskId, nil
}

func (h AddNewTaskToGroupHandler) createTaskAndSave(taskDto query.TaskDto) (domain.TaskID, error) {
	pr, err := domain.NewPriorityFromUint8(taskDto.Priority)
	if err != nil {
		return domain.TaskID{}, err
	}

	task := domain.NewTaskWithCurrentDate(taskDto.Title, taskDto.Description, pr)

	errTask := h.taskRepos.Save(task)
	if errTask != nil {
		return domain.TaskID{}, errTask
	}

	return task.ID(), nil
}

func (h AddNewTaskToGroupHandler) addNewTaskToGroup(groupUuid string, taskId domain.TaskID) error {
	group, err := h.getGroupByID(groupUuid)
	if err != nil {
		return err
	}

	task, errDb := h.getTaskByID(taskId)
	if errDb != nil {
		return errDb
	}

	errGroup := group.AddTask(task.ID())
	if errGroup != nil {
		return errGroup
	}

	errRepos := h.groupRepos.Save(group)
	if errRepos != nil {
		return errRepos
	}

	return nil
}

func (h AddNewTaskToGroupHandler) getGroupByID(groupUuid string) (*domain.Group, error) {
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

func (h AddNewTaskToGroupHandler) getTaskByID(taskId domain.TaskID) (*domain.Task, error) {
	task, err := h.taskRepos.GetByID(taskId)
	if err != nil {
		return &domain.Task{}, err
	}

	return task, nil
}
