package domain

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GroupID struct {
	value uuid.UUID
}

func (t GroupID) String() string {
	return t.value.String()
}

func NewGroupID() GroupID {
	return GroupID{uuid.New()}
}

func NewGroupIDFromString(v string) (GroupID, error) {
	value, err := uuid.Parse(v)
	return GroupID{value}, errors.Wrapf(err, "%s is invalid group ID format", v)
}

type Group struct {
	id          GroupID
	title       string
	description string
	taskIDs     []TaskID // TODO переделать на мапу
}

func NewGroup(title string, description string) *Group {
	return &Group{
		id:          NewGroupID(),
		title:       title,
		description: description,
	}
}

func Build(id GroupID, title string, description string, taskIDs []TaskID) *Group {
	return &Group{
		id:          id,
		title:       title,
		description: description,
		taskIDs:     taskIDs,
	}
}

func (g *Group) ID() GroupID {
	return g.id
}

func (g *Group) Tasks() []TaskID {
	return g.taskIDs
}

func (g *Group) Title() string {
	return g.title
}

func (g *Group) Description() string {
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
