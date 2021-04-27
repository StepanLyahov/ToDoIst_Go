package query

import (
	"domain"
	"time"
)

type GroupDto struct {
	id          string
	title       string
	description string
	taskIDs     []string
}

func GroupToDto(group domain.Group) GroupDto {
	taskIDStr := make([]string, 0)
	for _, id := range group.Tasks() {
		taskIDStr = append(taskIDStr, id.String())
	}

	dto := GroupDto{
		id:          group.ID().String(),
		title:       group.Title(),
		description: group.Description(),
		taskIDs:     taskIDStr,
	}

	return dto
}

type TaskDto struct {
	id               string
	title            string
	description      string
	priority         uint8
	createDate       time.Time
	currentDoingDate time.Time
	endDate          time.Time
}

func TaskToDto(task domain.Task) TaskDto {
	dto := TaskDto{
		id:               task.ID().String(),
		title:            task.Title(),
		description:      task.Description(),
		priority:         task.Priority().Uint8(),
		createDate:       task.CreateData(),
		currentDoingDate: task.CurrentData(),
		endDate:          task.EndDate(),
	}

	return dto
}
