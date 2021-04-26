package service

import (
	"app/repository"
	"domain"
)

type GroupService struct {
	groupRepos repository.GroupRepository
	taskRepos  repository.TaskRepository
}

func NewGroupService(groupRepos repository.GroupRepository,
	taskRepos repository.TaskRepository) *GroupService {

	return &GroupService{
		groupRepos: groupRepos,
		taskRepos: taskRepos,
	}
}

func (gs *GroupService) Create(title string, description string) (domain.GroupID, error) {
	group := domain.NewGroup(title, description)

	err := gs.groupRepos.Save(group)
	if err != nil {
		return domain.GroupID{}, err
	}

	return group.ID(), nil
}

func (gs *GroupService) Delete(uuid string) error {
	groupID, err := domain.NewGroupIDFromString(uuid)
	if err != nil {
		return err
	}
	errRepos := gs.groupRepos.DelByID(groupID)
	if errRepos != nil {
		return errRepos
	}

	return nil
}

func (gs *GroupService) AddNewTaskToGroup(groupUuid string, taskUuid string) error {
	group, err := gs.GetGroupByID(groupUuid)
	if err != nil {
		return err
	}

	task, errDb := gs.getTaskByID(taskUuid)
	if errDb != nil {
		return errDb
	}

	errGroup := group.AddTask(task.ID())
	if errGroup != nil {
		return errGroup
	}

	errRepos := gs.groupRepos.Save(group)
	if errRepos != nil {
		return errRepos
	}

	return nil
}

func (gs *GroupService) DelTaskFromGroup(groupUuid string, taskUuid string) error {
	group, err := gs.GetGroupByID(groupUuid)
	if err != nil {
		return err
	}

	task, errDb := gs.getTaskByID(taskUuid)
	if errDb != nil {
		return errDb
	}

	group.DelTask(task.ID())

	errRepos := gs.groupRepos.Save(group)
	if errRepos != nil {
		return errRepos
	}

	return nil
}

func (gs *GroupService) GetAll() []*domain.Group {
	return gs.groupRepos.GetAll()
}

func (gs *GroupService) GetGroupByID(groupUuid string) (*domain.Group, error) {
	groupID, err := domain.NewGroupIDFromString(groupUuid)
	if err != nil {
		return &domain.Group{}, err
	}

	group, err := gs.groupRepos.GetByID(groupID)
	if err != nil {
		return &domain.Group{}, err
	}

	return group, nil
}

func (gs *GroupService) getTaskByID(groupUuid string) (*domain.Task, error) {
	taskID, err := domain.NewTaskIDFromString(groupUuid)
	if err != nil {
		return &domain.Task{}, err
	}

	task, err := gs.taskRepos.GetByID(taskID)
	if err != nil {
		return &domain.Task{}, err
	}

	return task, nil
}
