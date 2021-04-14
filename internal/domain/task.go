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
	createDate  time.Time
	currentDate time.Time
	endDate     time.Time
}

type Task struct {
	title       string
	description string
	optionsDate optionsDate
	priority    Priority
}

func (t *Task) GetCreateData() time.Time {
	return t.optionsDate.createDate
}

func (t *Task) GetCurrentData() time.Time {
	return t.optionsDate.currentDate
}

func (t *Task) GetEndDate() time.Time {
	return t.optionsDate.endDate
}

func (t *Task) SetCreateData(d time.Time) {
	t.optionsDate.createDate = d
}

func (t *Task) SetCurrentData(d time.Time) {
	t.optionsDate.currentDate = d
}

func (t *Task) SetEndDate(d time.Time) {
	t.optionsDate.endDate = d
}

func (t *Task) SetLevelFrom1To5(priority Priority) {
	t.priority = priority
}
