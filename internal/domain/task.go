package domain

import (
	"github.com/pkg/errors"
	"time"
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
	level       uint8
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

func (t *Task) SetLevelFrom1To5(level uint8) error {
	if level < 1 && level > 5 {
		return errors.New("level must be [1..5]")
	}

	t.level = level
	return nil
}
