package command

import (
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository"
	"log"
	"testing"
)

func TestDeleteGroup(t *testing.T) {
	h := NewDeleteGroupHandler(repository.NewInMemoryGroup())

	group := domain.NewGroup("", "")
	log.Print(group.ID())

	err := h.groupRepos.Save(group)
	if err != nil {
		t.Fatalf("It shouldn't be a problem, but Err := %v", err)
	}

	err = h.Execute(group.ID().String())
	if err != nil {
		t.Fatalf("Execute musted be del by ID, but Err := %v", err)
	}

	_, err = h.groupRepos.GetByID(group.ID())
	if err == nil {
		t.Fatalf("Err must be 'not found' but err is nil")
	}

}

func TestDeleteGroupWithInvalidID(t *testing.T) {
	h := NewDeleteGroupHandler(repository.NewInMemoryGroup())

	err := h.Execute("invalid uuid")
	if err == nil {
		t.Fatalf("uuid must be invalid")
	}

}
