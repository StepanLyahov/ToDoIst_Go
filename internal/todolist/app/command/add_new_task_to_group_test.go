package command

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app/query"
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository"
	"testing"
)

func initGroupHandlerBeforeTest(h *AddNewTaskToGroupHandler) domain.GroupID {
	gDto := GroupDTO{
		Title:       "Title",
		Description: "Description",
	}

	group := domain.NewGroup(gDto.Title, gDto.Description)
	err := h.groupRepos.Save(group)
	if err != nil {
		panic(err)
	}

	return group.ID()
}

func TestNewAddNewTaskToGroup(t *testing.T) {
	h := NewAddNewTaskToGroupHandler(repository.NewInMemoryGroup(), repository.NewInMemoryTask())
	groupId := initGroupHandlerBeforeTest(&h)

	taskDto := query.TaskDto{
		Title:       "Title",
		Description: "Description",
		Priority:    2,
	}

	taskID, err := h.Execute(groupId.String(), taskDto)
	if err != nil {
		t.Fatalf("Err must be nil, but err %v", err)
	}

	task, err := h.taskRepos.GetByID(taskID)
	if err != nil {
		t.Fatalf("Err must be nil, but err %v", err)
	}

	if !compareTask(taskDto, *task) {
		t.Fatalf("Dto = '%v' and Group = '%v' does not compare",
			taskDto, *task)
	}
}

func compareTask(gDto query.TaskDto, task domain.Task) bool {
	if gDto.Description != task.Description() {
		return false
	}
	if gDto.Title != task.Title() {
		return false
	}
	return true
}
