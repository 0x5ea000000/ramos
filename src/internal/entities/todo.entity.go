package entities

type Todo struct {
	Name string `json:"name"`
	Note string `json:"note"`
}

var TodoList []Todo
