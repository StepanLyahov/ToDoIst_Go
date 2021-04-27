package domain

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

type Priority uint8

const (
	Priority1 Priority = iota + 1
	Priority2
	Priority3
	Priority4
	Priority5
	Error
)

func NewPriorityFromUint8(priority uint8) (Priority, error) {
	switch priority {
	case 1:
		return Priority1, nil
	case 2:
		return Priority2, nil
	case 3:
		return Priority3, nil
	case 4:
		return Priority4, nil
	case 5:
		return Priority5, nil
	default:
		return Error, errors.New("can't parse to Priority")
	}
}

type TaskID struct {
	value uuid.UUID
}

func NewTaskID() TaskID {
	return TaskID{uuid.New()}
}

func NewTaskIDFromString(v string) (TaskID, error) {
	value, err := uuid.Parse(v)
	return TaskID{value}, errors.Wrapf(err, "%s is invalid task ID format", v)
}

type optionsDate struct {
	createDate       time.Time
	currentDoingDate time.Time
	endDate          time.Time
}

type Task struct {
	id          TaskID
	title       string
	description string
	optionsDate optionsDate
	priority    Priority
}

func NewTaskWithCurrentDate(title string, description string, priority Priority) *Task {
	oDate := optionsDate{
		createDate:       time.Now(),
		currentDoingDate: time.Now(),
		endDate:          time.Now(),
	}

	return &Task{
		NewTaskID(),
		title,
		description,
		oDate,
		priority,
	}
}

func (t *Task) CreateData() time.Time {
	return t.optionsDate.createDate
}

func (t *Task) CurrentData() time.Time {
	return t.optionsDate.currentDoingDate
}

func (t *Task) EndDate() time.Time {
	return t.optionsDate.endDate
}

func (t *Task) SetCreateData(d time.Time) {
	t.optionsDate.createDate = d
}

func (t *Task) SetCurrentData(d time.Time) {
	t.optionsDate.currentDoingDate = d
}

func (t *Task) SetEndDate(d time.Time) {
	t.optionsDate.endDate = d
}

func (t *Task) SetLevelFrom1To5(priority Priority) {
	t.priority = priority
}

func (t *Task) ID() TaskID {
	return t.id
}
