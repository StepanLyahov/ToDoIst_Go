package query

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app/repository"
	"github.com/pkg/errors"
	"log"
)

type GetAllGroupHandler struct {
	groupRepos repository.GroupRepository
}

func NewGetAllGroupHandler(rep repository.GroupRepository) GetAllGroupHandler {
	return GetAllGroupHandler{groupRepos: rep}
}

func (h *GetAllGroupHandler) Execute() ([]GroupDto, error) {
	groups, err := h.groupRepos.GetAll()
	if err != nil {
		log.Printf("Error get all Groups: %v", err)
		return nil, errors.New("ошибка получения всех групп")
	}

	groupDtos := make([]GroupDto, 0)

	for _, group := range groups {
		dto := GroupToDto(*group)
		groupDtos = append(groupDtos, dto)
	}
	return groupDtos, nil
}
