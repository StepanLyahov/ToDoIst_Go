package domain

import (
	"errors"
	"github.com/google/uuid"
)

type GroupID struct {
	value uuid.UUID
}

func NewGroupID() GroupID {
	return GroupID{uuid.New()}
}

type Group struct {
	id          GroupID
	title       string
	description string
	taskIDs     []TaskID
}

func NewGroup(title string, description string) *Group {
	return &Group{
		id:          NewGroupID(),
		title:       title,
		description: description,
	}
}

func (g *Group) ID() GroupID {
	return g.id
}

func (g *Group) GetTasks() []TaskID {
	return g.taskIDs
}

func (g *Group) GetTitle() string {
	return g.title
}

func (g *Group) GetDescription() string {
	return g.description
}

func (g *Group) AddTask(id TaskID) error {
	if contains(g.taskIDs, id) {
		return errors.New("Task with current id is already exists")
	}

	g.taskIDs = append(g.taskIDs, id)
	return nil
}

func (g *Group) DelTask(id TaskID) {
	for pos, taskId := range g.taskIDs {
		if id == taskId {
			g.taskIDs = removeIndex(g.taskIDs, pos)
		}
	}
}

func removeIndex(s []TaskID, index int) []TaskID {
	return append(s[:index], s[index+1:]...)
}

func contains(ids []TaskID, id TaskID) bool {
	for _, a := range ids {
		if a == id {
			return true
		}
	}
	return false
}
