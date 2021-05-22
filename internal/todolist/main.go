package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"pkg/server"
	"todolist/app"
	"todolist/app/command"
	"todolist/app/query"
	"todolist/infrastructure/repository"
	"todolist/infrastructure/web"
)

func main() {
	app := newApplication()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return web.HandlerFromMux(web.NewHTTPServer(app), router)
	})
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
