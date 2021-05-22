package command

import (
	"todolist/app/repository"
	"todolist/domain"
)

type DeleteGroupHandler struct {
	groupRepos repository.GroupRepository
}

func NewDeleteGroupHandler(rep repository.GroupRepository) DeleteGroupHandler {
	return DeleteGroupHandler{groupRepos: rep}
}

func (h DeleteGroupHandler) Execute(uuid string) error {
	groupID, err := domain.NewGroupIDFromString(uuid)
	if err != nil {
		return err
	}
	errRepos := h.groupRepos.DelByID(groupID)
	if errRepos != nil {
		return errRepos
	}

	return nil
}
