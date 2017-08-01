package main

import (
	"fmt"
)

var	currentId	int
var todos		ToDos

func init() {
	RepoCreateTodo(ToDo{Name: "Hello, Execute Task #001"})
	RepoCreateTodo(ToDo{Name: "Hello, Execute Task #002"})
}

func RepoFundTodo(id int) ToDo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return ToDo{}
}

func RepoCreateTodo(t ToDo) ToDo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyTodo(id int) error {
    for i, t := range todos {
        if t.Id == id {
            todos = append(todos[:i], todos[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}

