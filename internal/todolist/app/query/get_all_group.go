package query

import "app/app/repository"

type GetAllGroupHandler struct {
	groupRepos repository.GroupRepository
}

func NewGetAllGroupHandler(rep repository.GroupRepository) GetAllGroupHandler {
	return GetAllGroupHandler{groupRepos: rep}
}

func (h *GetAllGroupHandler) Execute() []GroupDto {
	groups := h.groupRepos.GetAll()

	groupDtos := make([]GroupDto, 0)

	for _, group := range groups {
		dto := GroupToDto(*group)
		groupDtos = append(groupDtos, dto)
	}
	return groupDtos
}
