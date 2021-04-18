package domain

import "testing"

func TestGroup_AddTaskIDInVoidGroup(t *testing.T) {
	taskID1 := NewTaskID()

	group := NewGroup("title", "description")
	_ = group.AddTask(taskID1)

	if isContainsTaskId(group.taskIDs, taskID1) != true {
		t.Fatalf("AddTask did't add new taskID")
	}
}

func TestGroup_AddDubbedTask(t *testing.T) {
	taskID1 := NewTaskID()

	group := NewGroup("title", "description")
	_ = group.AddTask(taskID1)
	err := group.AddTask(taskID1)

	if err == nil {
		t.Fatalf("Err must be not nil")
	}
}

func TestGroup_DelTaskID(t *testing.T) {
	taskID1 := NewTaskID()

	group := NewGroup("title", "description")
	_ = group.AddTask(taskID1)

	group.DelTask(taskID1)

	if isContainsTaskId(group.taskIDs, taskID1) != false {
		t.Fatalf("DelTask did't del taskID(%v) from group.", taskID1)
	}
}

func isContainsTaskId(ids []TaskID, id TaskID) bool {
	for _, taskID := range ids {
		if taskID == id {
			return true
		}
	}
	return false
}
