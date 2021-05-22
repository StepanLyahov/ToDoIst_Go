package web

import (
	"encoding/json"
	"fmt"
	"github.com/StepanLyahov/ToDoIst/todolist/app"
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
		panic(err)
	}


	res := CreateGroupRequest{}
	err = json.Unmarshal(reqBody, &res)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)
}
