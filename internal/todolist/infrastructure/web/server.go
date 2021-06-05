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
		responseWithErr("Err read body", err.Error(), w, 500)
		return
	}

	res := CreateGroupRequest{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		responseWithErr("Err parse body", err.Error(), w, 500)
		return
	}

	groupDto := command.GroupDTO{
		Title:       res.Title,
		Description: res.Description,
	}

	execute, err := H.app.Commands.CreateGroup.Execute(groupDto)
	if err != nil {
		responseWithErr("Internal err", err.Error(), w, 500)
		return
	}

	w.Header().Set("Content-Location", fmt.Sprintf("/group/%s", execute))
	w.WriteHeader(http.StatusCreated)
}

func responseWithErr(slug string, details string, w http.ResponseWriter, status int) {
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
