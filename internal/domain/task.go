package domain

import (
	"time"
)

type Priority uint8

const (
	Priority1 Priority = iota + 1
	Priority2
	Priority3
	Priority4
	Priority5
)

type optionsDate struct {
	createDate       time.Time
	currentDoingDate time.Time
	endDate          time.Time
}

type Task struct {
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
		title:       title,
		description: description,
		optionsDate: oDate,
		priority:    priority,
	}
}

func (t *Task) GetCreateData() time.Time {
	return t.optionsDate.createDate
}

func (t *Task) GetCurrentData() time.Time {
	return t.optionsDate.currentDoingDate
}

func (t *Task) GetEndDate() time.Time {
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
