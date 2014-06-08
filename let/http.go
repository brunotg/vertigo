package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/vmarmol/vertigo/let/api"
)

type taskExport struct {
	taskManager TaskManager
}

func (self *taskExport) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	args := strings.Split(r.URL.Path, "/")
	switch strings.ToUpper(r.Method) {
	case "GET":
		id := args[len(args)-1]
		fmt.Printf("export %v\n", id)
		c := &api.ContainerSpec{
			Id: id,
		}
		self.taskManager.Export(c, w)
	}
}

type restTaskManager struct {
	taskManager TaskManager
}

func (self *restTaskManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch strings.ToUpper(r.Method) {
	case "POST":
		decoder := json.NewDecoder(r.Body)
		var runspec api.RunSpec
		err := decoder.Decode(&runspec)
		if err != nil {
			fmt.Fprintf(w, "untable to decode: %v", err)
		}
		c, err := self.taskManager.RunTask(&runspec)
		if err != nil {
			fmt.Fprintf(w, "untable to run: %v", err)
		}
		encoder := json.NewEncoder(w)
		encoder.Encode(c)
	}
}
