package main

import (
	_"fmt"
	_"log"
	"net/http"
	_"github.com/gorilla/mux"
)

type Route struct {
	Name			string
	Method			string
	Pattern			string
	HandleFunc		http.HandlerFunc
}

type Routes []Route

var routes = Routes {
	Route{
		"index",
		"GET",
		"/",
		index,
	},
	Route{
		"todoIndex",
		"GET",
		"/todo",
		todoIndex,
	},
	Route{
        "TodoShow",
        "GET",
        "/todo/{doId}",
        todoAction,
    },
	Route{
		"ToDoCreate",
		"GET",
		"/todos",
		TodoCreate,
	},
}
