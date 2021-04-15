package domain

import (
	"fmt"
)

type Group struct {
	title       string
	description string
	tasks       []Task
}

func (g *Group) getTasks() []Task {
	return g.tasks
}

func (g *Group) getTitle() string {
	return g.title
}

func (g *Group) getDescription() string {
	return g.description
}

func (g *Group) findTaskByTitle(title string) (Task, error) {

	for _, task := range g.tasks {
		if task.title == title {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("Task by title %v is not found", title)
}
