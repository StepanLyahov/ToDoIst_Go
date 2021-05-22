package app

import (
	"github.com/StepanLyahov/ToDoIst/todolist/app/command"
	"github.com/StepanLyahov/ToDoIst/todolist/app/query"
)

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	AddNewTaskToGroup   command.AddNewTaskToGroupHandler
	ChangeTask          command.ChangeTaskHandler
	CreateGroup         command.CreateGroupHandler
	DeleteGroup         command.DeleteGroupHandler
	DeleteTaskFromGroup command.DeleteTaskFromGroupHandler
}

type Queries struct {
	GetAllGroup  query.GetAllGroupHandler
	GetGroupById query.GetGroupByIdHandler
	GetTaskById  query.GetTaskByIdHandler
}
