package main

import (
	"time"	
)

type ToDo struct {
    Id			int			`json: "id"`
	Name        string      `json: "name"`
    Complete    bool        `json: "complete"`
    Due         time.Time   `json: "due"`
}

type ToDos []ToDo
