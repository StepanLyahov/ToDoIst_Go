package web

import (
	"encoding/json"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/app"
	"github.com/StepanLyahov/ToDoIst/todolist/app/command"
	"io/ioutil"
	"net/http"
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
		buildResponseWithErr("Err read body", err.Error(), w, 500)
		return
	}

	res := CreateGroupRequest{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		buildResponseWithErr("Err parse body", err.Error(), w, 500)
		return
	}

	groupDto := command.CreateGroupDTO{
		Title:       res.Title,
		Description: res.Description,
	}

	execute, err := H.app.Commands.CreateGroup.Execute(groupDto)
	if err != nil {
		buildResponseWithErr("Internal err", err.Error(), w, 500)
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
