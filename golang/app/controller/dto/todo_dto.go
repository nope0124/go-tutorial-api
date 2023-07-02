package dto

type FetchTodoResponse struct {
    Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type AddTodoRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type ChangeTodoRequest struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type DeleteTodoRequest struct {
	Id string `json:"id"`
}
