package main

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app"
	"github.com/StepanLyahov/ToDoIst/todolist/app/command"
	"github.com/StepanLyahov/ToDoIst/todolist/app/query"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository/in_memory"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/web"
	"github.com/StepanLyahov/ToDoIst/todolist/support/server"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {
	application := newApplication()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return web.HandlerFromMux(web.NewHTTPServer(application), router)
	})
}

func newApplication() app.Application {
	taskRep := in_memory.NewInMemoryTask()
	groupRep := in_memory.NewInMemoryGroup()

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
