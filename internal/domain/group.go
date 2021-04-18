package domain

type Group struct {
	title       string
	description string
	taskIDs     []TaskID
}

func (g *Group) getTasks() []TaskID {
	return g.taskIDs
}

func (g *Group) getTitle() string {
	return g.title
}

func (g *Group) getDescription() string {
	return g.description
}

func (g *Group) AddTask(id TaskID) {
	g.taskIDs = append(g.taskIDs, id)
}
