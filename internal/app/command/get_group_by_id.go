package command

import (
	"app/repository"
	"domain"
)

type GetGroupByIdHandler struct {
	groupRepos repository.GroupRepository
}

func NewGetGroupByIdHandler(rep repository.GroupRepository) GetGroupByIdHandler {
	return GetGroupByIdHandler{groupRepos: rep}
}

func (h GetGroupByIdHandler) Execute(groupUuid string) (*domain.Group, error) {
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