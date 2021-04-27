package command

import (
	"app/repository"
	"domain"
)

type GetAllGroupHandler struct {
	groupRepos repository.GroupRepository
}

func NewGetAllGroupHandler(rep repository.GroupRepository) GetAllGroupHandler {
	return GetAllGroupHandler{groupRepos: rep}
}

func (h *GetAllGroupHandler) Execute() []*domain.Group {
	return h.groupRepos.GetAll()
}