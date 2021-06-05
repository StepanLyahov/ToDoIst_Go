package query

import (
	"github.com/StepanLyahov/ToDoIst/todolist/domain"
	"time"
)

type GroupDto struct {
	Id          string
	Title       string
	Description string
	TaskIDs     []string
}

func GroupToDto(group domain.Group) GroupDto {
	taskIDStr := make([]string, 0)
	for _, id := range group.Tasks() {
		taskIDStr = append(taskIDStr, id.String())
	}

	dto := GroupDto{
		Id:          group.ID().String(),
		Title:       group.Title(),
		Description: group.Description(),
		TaskIDs:     taskIDStr,
	}

	return dto
}

type TaskDto struct {
	Id               string
	Title            string
	Description      string
	Priority         uint8
	CreateDate       time.Time
	CurrentDoingDate time.Time
	EndDate          time.Time
}

func TaskToDto(task domain.Task) TaskDto {
	dto := TaskDto{
		Id:               task.ID().String(),
		Title:            task.Title(),
		Description:      task.Description(),
		Priority:         task.Priority().Uint8(),
		CreateDate:       task.CreateData(),
		CurrentDoingDate: task.CurrentData(),
		EndDate:          task.EndDate(),
	}

	return dto
}

func TaskToDomain(dto TaskDto) (domain.Task, error) {
	d := domain.Task{}

	d.SetDescription(dto.Description)

	uuid, err := domain.NewTaskIDFromString(dto.Id)
	if err != nil {
		return domain.Task{}, err
	}

	d.SetId(uuid)
	d.SetTitle(dto.Title)

	d.SetCreateData(dto.CreateDate)
	d.SetCurrentData(dto.CurrentDoingDate)
	d.SetEndDate(dto.EndDate)

	pr, err := domain.NewPriorityFromUint8(dto.Priority)
	if err != nil {
		return domain.Task{}, err
	}
	d.SetPriority(pr)

	return d, nil
}
