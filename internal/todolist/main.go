package main

import (
	"database/sql"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/app"
	"github.com/StepanLyahov/ToDoIst/todolist/app/command"
	"github.com/StepanLyahov/ToDoIst/todolist/app/query"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository/in_memory"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/repository/postgres"
	"github.com/StepanLyahov/ToDoIst/todolist/infrastructure/web"
	"github.com/StepanLyahov/ToDoIst/todolist/support/server"
	"github.com/go-chi/chi/v5"
	"github.com/labstack/gommon/log"
	"net/http"
	"sync"
)

func main() {
	application := newApplication()

	server.RunHTTPServer(func(router chi.Router) http.Handler {
		return web.HandlerFromMux(web.NewHTTPServer(application), router)
	})
}

func newApplication() app.Application {
	db, err := initPostgresConnection()
	if err != nil {
		panic(err)
	}

	taskRep := in_memory.NewInMemoryTask()
	groupRep := postgres.NewPostgresGroup(db)

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

var once sync.Once

func initPostgresConnection() (*sql.DB, error) {
	user := "postgres"
	password := "postgres"
	dbname := "stepanlahov"
	dbtype := "postgres"

	var err error
	var db *sql.DB

	once.Do(func() {
		connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable",
			user, password, dbname)

		db, err = sql.Open(dbtype, connStr)

		log.Print("Connected!!!")
	})

	return db, err
}
