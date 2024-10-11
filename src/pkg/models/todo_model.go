package models

type Todo struct {
	BaseModel
	Name string `json:"name"`
	Note string `json:"note"`
}
