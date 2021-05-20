package command

import (
	"app/domain"
	"app/infrastructure/repository"
	"testing"
)

func TestDeleteTaskFromGroupValidID(t *testing.T) {
	repGroup := repository.NewInMemoryGroup()
	repTask := repository.NewInMemoryTask()

	taskID, groupID := addRelatedGroupAndTaskInRepo(repTask, repGroup)

	h := NewDeleteTaskFromGroupHandler(repGroup, repTask)

	err := h.Execute(groupID.String(), taskID.String())

	if err != nil {
		t.Fatalf("Err must be nil, but now err = %v", err)
	}

	group, err := repGroup.GetByID(groupID)
	if err != nil {
		t.Fatalf("Group Repo must have group by id{%v}, but err = %v", groupID, err)
	}

	if findById(group.Tasks(), taskID) == true {
		t.Fatalf("After doing 'Execute' group does not have the taskID")
	}
}

func addRelatedGroupAndTaskInRepo(tRep repository.InMemoryTask, gRep repository.InMemoryGroup) (taskID domain.TaskID, groupID domain.GroupID) {

	task := domain.NewTaskWithCurrentDate("title", "desc", domain.Priority1)
	group := domain.NewGroup("titleG", "descG")

	_ = group.AddTask(task.ID())

	_ = tRep.Save(task)
	_ = gRep.Save(group)

	return task.ID(), group.ID()
}

func findById(list []domain.TaskID, val domain.TaskID) bool {
	for _, id := range list {
		if id == val {
			return true
		}
	}
	return false
}
