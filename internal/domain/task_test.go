package domain

import "testing"

func TestTask_SetLevelFrom1To5_SetInvalidLevel(t *testing.T) {
	task := Task{}

	err := task.SetLevelFrom1To5(6)

	if err == nil {
		t.Fatalf("Err must be `level must be [1..5]` but now is nil")
	}
}
