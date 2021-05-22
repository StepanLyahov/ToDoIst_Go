package command

import (
	"testing"
	"todolist/app/query"
	"todolist/domain"
	"todolist/infrastructure/repository"
)

func TestChangeTask(t *testing.T) {
	title := "lol"

	h := NewChangeTaskHandler(repository.NewInMemoryTask())
	task := initTaskBeforeTest(&h)

	dtoTask := query.TaskToDto(task)
	dtoTask.Title = title

	err := h.Execute(dtoTask)
	if err != nil {
		t.Fatalf("Execute must be ok, but err := %v", err)
	}

	updatedTask, err := h.taskRepos.GetByID(task.ID())
	if err != nil {
		t.Fatalf("In repository must be task by id := %v, but err := %v", task.ID().String(), err)
	}

	if updatedTask.Title() != title {
		t.Fatalf("Title in updated task{%v} must be %v", updatedTask, title)
	}

}

func initTaskBeforeTest(h *ChangeTaskHandler) domain.Task {
	task := domain.NewTaskWithCurrentDate("t", "d", domain.Priority1)
	err := h.taskRepos.Save(task)
	if err != nil {
		panic(err)
	}

	return *task
}
