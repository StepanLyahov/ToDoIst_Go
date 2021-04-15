package domain

import (
	"fmt"
)

type Gruop struct {
	title       string
	description string
	tasks       []Task
}

func (g *Gruop) getTasks() []Task {
	return g.tasks
}

func (g *Gruop) getTitle() string {
	return g.title
}

func (g *Gruop) getDescription() string {
	return g.description
}

func (g *Gruop) findTaskByTitle(title string) (Task, error) {

	for _, task := range g.tasks {
		if task.title == title {
			return task, nil
		}
	}

	return Task{}, fmt.Errorf("Task by title %v is not found", title)
}
