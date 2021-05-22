package command

import (
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository"
	"testing"
)

func TestCreateGroup(t *testing.T) {
	gDto := GroupDTO{
		Title:       "Title",
		Description: "Description",
	}

	h := NewCreateGroupHandler(repository.NewInMemoryGroup())
	id, err := h.Execute(gDto)

	if err != nil {
		t.Fatalf("Expected what 'Execute' return (id,nil),"+
			" but now Err := %v", err)
	}

	group, err := h.groupRepos.GetByID(id)

	if err != nil {
		t.Fatalf("Expected what repository contain Group,"+
			" but err := %v", err)
	}

	if !compare(gDto, *group) {
		t.Fatalf("Dto = '%v' and Group = '%v' does not compare",
			gDto, *group)
	}

}
func compare(gDto GroupDTO, group domain.Group) bool {
	if gDto.Description != group.Description() {
		return false
	}
	if gDto.Title != group.Title() {
		return false
	}

	return true
}
