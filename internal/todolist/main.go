package main

import (
	"app/app"
	"app/app/command"
	"app/app/query"
	"app/infrastructure/repository"
)

func main() {

}

func newApplication() app.Application {
	taskRep := repository.NewInMemoryTask()
	groupRep := repository.NewInMemoryGroup()

	com := app.Commands{
		AddNewTaskToGroup:   command.NewAddNewTaskToGroupHandler(groupRep, taskRep),
		ChangeTask:          command.NewChangeTaskHandler(taskRep),
		CreateGroup:         command.NewCreateGroupHandler(groupRep),
		DeleteGroup:         command.NewDeleteGroupHandler(groupRep),
		DeleteTaskFromGroup: command.NewDeleteTaskFromGroupHandler(groupRep, taskRep),
	}

	que := app.Queries{
		GetAllGroup:  query.NewGetAllGroupHandler(groupRep),
		GetGroupById: query.NewGetGroupByIdHandler(groupRep),
		GetTaskById:  query.NewGetTaskByIdHandler(taskRep),
	}

	return app.Application{
		Commands: com,
		Queries:  que,
	}
}
