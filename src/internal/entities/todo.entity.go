package entities

type Todo struct {
	ID   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Note string `json:"note"`
}
