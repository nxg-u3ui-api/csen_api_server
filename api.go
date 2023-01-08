package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type api struct {
	group *group
}

type Message struct {
	Id, Seq, Caller, Callee string
}

func (api *api) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var msg Message
	err := json.NewDecoder(req.Body).Decode(&msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("%+v\n", msg)

	data, err := json.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	api.group.bridge <- data
}
