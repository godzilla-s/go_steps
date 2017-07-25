package main

import (
	"fmt"
	"log"
	"io"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Heelo, Welcome to Index!!")
}

func todoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//设置状态码
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }	
}

func todoAction(w http.ResponseWriter, r *http.Request) {
	args := mux.Vars(r)
	log.Println(args)
	doId := args["doId"]
	fmt.Fprintln(w, "Action: ", doId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
    var todo ToDo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }

    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

    t := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

