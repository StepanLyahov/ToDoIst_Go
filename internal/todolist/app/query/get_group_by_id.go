package query

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app/repository"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
)

type GetGroupByIdHandler struct {
	groupRepos repository.GroupRepository
}

func NewGetGroupByIdHandler(rep repository.GroupRepository) GetGroupByIdHandler {
	return GetGroupByIdHandler{groupRepos: rep}
}

func (h GetGroupByIdHandler) Execute(groupUuid string) (GroupDto, error) {
	groupID, err := domain.NewGroupIDFromString(groupUuid)
	if err != nil {
		return GroupDto{}, err
	}

	group, err := h.groupRepos.GetByID(groupID)
	if err != nil {
		return GroupDto{}, err
	}
	groupDtp := GroupToDto(*group)

	return groupDtp, nil
}
