package web

import (
	"encoding/json"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/app"
	"github.com/StepanLyahov/ToDoIst/todolist/app/command"
	"github.com/StepanLyahov/ToDoIst/todolist/app/query"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPServer struct {
	app app.Application
}

func NewHTTPServer(app app.Application) HTTPServer {
	return HTTPServer{app}
}

func (H HTTPServer) CreateGroup(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		buildResponseWithErr("Err read body", err.Error(), w, http.StatusInternalServerError)
		return
	}

	res := CreateGroupRequest{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		buildResponseWithErr("Err parse body", err.Error(), w, http.StatusInternalServerError)
		return
	}

	groupDto := command.CreateGroupDTO{
		Title:       res.Title,
		Description: res.Description,
	}

	execute, err := H.app.Commands.CreateGroup.Execute(groupDto)
	if err != nil {
		buildResponseWithErr("Internal err", err.Error(), w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Location", fmt.Sprintf("/group/%s", execute))
	w.WriteHeader(http.StatusCreated)
}

func (H HTTPServer) GetGroup(w http.ResponseWriter, r *http.Request, groupId string) {

	execute, err := H.app.Queries.GetGroupById.Execute(groupId)
	if err != nil {
		buildResponseWithErr("Err GetGroup by Id", err.Error(), w, http.StatusNotFound)
		return
	}

	response := GroupResponse{
		Description: &execute.Description,
		Id:          &execute.Id,
		TaskIDs:     &execute.TaskIDs,
		Title:       &execute.Title,
	}

	buildResponse(response, w, http.StatusCreated)
}

func (H HTTPServer) AddTaskInGroup(w http.ResponseWriter, r *http.Request, groupId string) {
	taskDto, err := bodyToTask(w, r)
	if err != nil {
		return
	}

	execute, err := H.app.Commands.AddNewTaskToGroup.Execute(groupId, taskDto)
	if err != nil {
		buildResponseWithErr("Err add task in group", err.Error(), w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Location", execute.String())
	w.WriteHeader(http.StatusCreated)
}

func (H HTTPServer) ChangeTask(w http.ResponseWriter, r *http.Request) {
	taskDto, err := bodyToTask(w, r)
	if err != nil {
		return
	}

	err = H.app.Commands.ChangeTask.Execute(taskDto)
	if err != nil {
		buildResponseWithErr("Err change task", err.Error(), w, http.StatusInternalServerError)
	}
}

func bodyToTask(w http.ResponseWriter, r *http.Request) (query.TaskDto, error) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		buildResponseWithErr("Err read body", err.Error(), w, http.StatusInternalServerError)
		return query.TaskDto{}, err
	}

	res := Task{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		buildResponseWithErr("Err parse body", err.Error(), w, http.StatusInternalServerError)
		return query.TaskDto{}, err
	}

	createDateParsed, err := time.Parse(time.RFC3339, *res.CreateDate)
	if err != nil {
		buildResponseWithErr("Err parse date", err.Error(), w, http.StatusInternalServerError)
		return query.TaskDto{}, err
	}

	currentDoingDateParsed, err := time.Parse(time.RFC3339, *res.CurrentDoingDate)
	if err != nil {
		buildResponseWithErr("Err parse date", err.Error(), w, http.StatusInternalServerError)
		return query.TaskDto{}, err
	}

	endDateParsed, err := time.Parse(time.RFC3339, *res.EndDate)
	if err != nil {
		buildResponseWithErr("Err parse date", err.Error(), w, http.StatusInternalServerError)
		return query.TaskDto{}, err
	}

	taskDto := query.TaskDto{
		Id:               *res.Id,
		Title:            *res.Title,
		Description:      *res.Description,
		Priority:         uint8(*res.Priority),
		CreateDate:       createDateParsed,
		CurrentDoingDate: currentDoingDateParsed,
		EndDate:          endDateParsed,
	}

	return taskDto, nil
}

func buildResponse(v interface{}, w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

func buildResponseWithErr(slug string, details string, w http.ResponseWriter, status int) {
	w.WriteHeader(status)

	errResponse := Error{
		details,
		slug,
	}

	err := json.NewEncoder(w).Encode(errResponse)
	if err != nil {
		panic(err)
	}
}
