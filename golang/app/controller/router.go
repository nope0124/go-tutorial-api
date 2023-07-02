package controller

import (
	"net/http"
	"os"
)

type Router interface {
	HandleRequest()
}

type router struct {
	tc TodoController
}

func CreateRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) HandleRequest() {
	http.HandleFunc("/todo/", ro.HandleTodoRequest)
	http.HandleFunc("/todos", ro.HandleTodosRequest)
}

func (ro *router) HandleTodoRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", os.Getenv("ORIGIN"))
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "OPTIONS" {
		return
	}

	switch r.URL.Path {
	case "/todos":
        if r.Method == "GET" {
            ro.tc.FetchTodos(w, r)
        } else if r.Method == "POST" {
			ro.tc.AddTodo(w, r)
		} else {
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    case "/todo/{id}":
        if r.Method == "DELETE" {
            ro.tc.DeleteTodo(w, r)
        } else if r.Method == "PUT" {
			ro.tc.ChangeTodo(w, r)
		} else {
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func (ro *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ro.tc.FetchTodos(w, r)
	} else if r.Method == "POST" {
		ro.tc.AddTodo(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}