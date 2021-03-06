package command

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app/repository"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
)

type CreateGroupHandler struct {
	groupRepos repository.GroupRepository
}

type CreateGroupDTO struct {
	Title       string
	Description string
}

func NewCreateGroupHandler(rep repository.GroupRepository) CreateGroupHandler {
	return CreateGroupHandler{groupRepos: rep}
}

func (h CreateGroupHandler) Execute(g CreateGroupDTO) (domain.GroupID, error) {
	group := domain.NewGroup(g.Title, g.Description)

	err := h.groupRepos.Save(group)
	if err != nil {
		return domain.GroupID{}, err
	}

	return group.ID(), nil
}
