package domain

type Gruop struct {
	title       string
	description string
	tasks       []Task
}

func (l *Gruop) getTasks() []Task {
	return l.tasks
}

func (l *Gruop) getTitle() string {
	return l.title
}

func (l *Gruop) getDescription() string {
	return l.description
}
